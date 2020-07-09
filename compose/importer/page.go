package importer

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/deinterfacer"
	"github.com/cortezaproject/corteza-server/pkg/importer"
)

type (
	Page struct {
		imp *Importer

		namespace *types.Namespace
		set       types.PageSet
		dirty     map[uint64]bool

		// page => module handle
		modules map[string]string

		// child => parent handle
		tree map[string][]string
	}

	// @todo remove finder strategy, directly provide set of items
	pageFinder interface {
		Find(filter types.PageFilter) (set types.PageSet, f types.PageFilter, err error)
	}
)

func NewPageImporter(imp *Importer, ns *types.Namespace) *Page {
	out := &Page{
		imp: imp,

		namespace: ns,

		set:   types.PageSet{},
		dirty: make(map[uint64]bool),

		modules: map[string]string{},
		tree:    map[string][]string{},
	}

	if imp.pageFinder != nil && ns.ID > 0 {
		out.set, _, _ = imp.pageFinder.Find(types.PageFilter{NamespaceID: ns.ID})
	}

	return out
}

func (pImp *Page) getModule(handle string) (*types.Module, error) {
	if g, ok := pImp.imp.namespaces.modules[pImp.namespace.Slug]; !ok {
		return nil, errors.Errorf("could not get modules %q from non existing namespace %q", handle, pImp.namespace.Slug)
	} else {
		return g.Get(handle)
	}
}

func (pImp *Page) getChart(handle string) (*types.Chart, error) {
	if g, ok := pImp.imp.namespaces.charts[pImp.namespace.Slug]; !ok {
		return nil, errors.Errorf("could not get chart %q from non existing namespace %q", handle, pImp.namespace.Slug)
	} else {
		return g.Get(handle)
	}
}

// CastSet Resolves permission rules:
// { <page-handle>: { page } } or [ { page }, ... ]
func (pImp *Page) CastSet(set interface{}) error {
	return pImp.castSet("", set)
}

// CastSet Resolves permission rules:
// { <page-handle>: { page } } or [ { page }, ... ]
func (pImp *Page) castSet(parent string, set interface{}) error {
	return deinterfacer.Each(set, func(index int, handle string, def interface{}) error {
		if index > -1 {
			// Pages defined as collection
			deinterfacer.KVsetString(&handle, "handle", def)
		}

		return pImp.cast(parent, handle, def)
	})
}

func (pImp *Page) Cast(handle string, def interface{}) (err error) {
	return pImp.cast("", handle, def)
}

// Cast Resolves permission rules:
// { <page-handle>: { page } } or [ { page }, ... ]
func (pImp *Page) cast(parent, handle string, def interface{}) (err error) {
	var page *types.Page

	if !importer.IsValidHandle(handle) {
		return errors.New("invalid page handle")
	}

	handle = importer.NormalizeHandle(handle)

	if page, err = pImp.Get(handle); err != nil {
		return err
	} else if page == nil {
		page = &types.Page{
			Handle:  handle,
			Title:   handle,
			Visible: true,
		}

		pImp.set = append(pImp.set, page)
	} else if page.ID == 0 {
		return errors.Errorf("page handle %q already defined in this import session", page.Handle)
	} else {
		pImp.dirty[page.ID] = true
	}

	pImp.tree[parent] = append(pImp.tree[parent], handle)

	if title, ok := def.(string); ok && title != "" {
		page.Title = title
		return nil
	}

	return deinterfacer.Each(def, func(_ int, key string, val interface{}) (err error) {
		switch key {
		case "namespace":
			// namespace value sanity check
			if deinterfacer.ToString(val, pImp.namespace.Slug) != pImp.namespace.Slug {
				return fmt.Errorf("explicitly set namespace on page %q shadows inherited namespace", pImp.namespace.Slug)
			}

		case "handle":
			// handle value sanity check
			if deinterfacer.ToString(val, handle) != handle {
				return fmt.Errorf("explicitly set handle on page %q shadows inherited handle", handle)
			}

		case "module":
			pImp.modules[handle] = deinterfacer.ToString(val)

		case "visible":
			page.Visible = deinterfacer.ToBool(val)

		case "title", "name", "label":
			page.Title = deinterfacer.ToString(val)

		case "description":
			page.Description = deinterfacer.ToString(val)

		case "blocks":
			return pImp.castBlocks(page, val)

		case "pages":
			return pImp.castSet(handle, val)

		case "allow", "deny":
			return pImp.imp.permissions.CastSet(types.PagePermissionResource.String()+handle, key, val)

		default:
			return fmt.Errorf("unexpected key %q for page %q", key, handle)
		}

		return err
	})
}

func (pImp *Page) castBlocks(page *types.Page, def interface{}) error {
	page.Blocks = types.PageBlocks{}

	return deinterfacer.Each(def, func(b int, _ string, blockDef interface{}) (err error) {
		block := types.PageBlock{}

		err = deinterfacer.Each(blockDef, func(_ int, key string, val interface{}) (err error) {
			switch key {
			case "title", "name", "label":
				block.Title = deinterfacer.ToString(val)

			case "description":
				block.Description = deinterfacer.ToString(val)

			case "kind":
				block.Kind = deinterfacer.ToString(val)

			case "options":
				block.Options, err = pImp.castBlockOptions(val)
				return err

			case "style":
				block.Style, err = pImp.castBlockStyle(page, b, val)
				return

			case "XYWH", "xywh", "dim", "dimension":
				xywh := deinterfacer.ToInts(val)
				if len(xywh) != 4 {
					return errors.New("invalid dimension (xywh) value, expecting slice with 4 integers")
				}

				block.XYWH = [4]int{xywh[0], xywh[1], xywh[2], xywh[3]}

			default:
				return fmt.Errorf("unexpected key %q for block on page %q", key, page.Handle)

			}

			return nil
		})

		if err != nil {
			return err
		}

		page.Blocks = append(page.Blocks, pImp.sanitizeBlock(block))
		return
	})
}

func (pImp *Page) castBlockOptions(def interface{}) (opt map[string]interface{}, err error) {
	opt = make(map[string]interface{})

	return opt, deinterfacer.Each(def, func(_ int, key string, val interface{}) (err error) {
		opt[key] = deinterfacer.Simplify(val)
		return nil
	})
}

func (pImp *Page) castBlockStyle(page *types.Page, n int, def interface{}) (s types.PageBlockStyle, err error) {
	s = types.PageBlockStyle{}

	return s, deinterfacer.Each(def, func(_ int, key string, val interface{}) (err error) {
		switch key {
		case "variants":
			s.Variants = map[string]string{}
			return deinterfacer.Each(val, func(_ int, key string, val interface{}) (err error) {
				s.Variants[key] = deinterfacer.ToString(val)
				return
			})
		default:
			return fmt.Errorf("unexpected key %q for block #%d on page %q", key, n+1, page.Handle)

		}
	})
}

func (pImp *Page) sanitizeBlock(b types.PageBlock) types.PageBlock {
	switch strings.ToLower(b.Kind) {
	case "automation":
		return pImp.sanitizeAutomationBlock(b)
	}

	return b
}

func (pImp *Page) sanitizeAutomationBlock(b types.PageBlock) types.PageBlock {
	b.Options["sealed"] = deinterfacer.ToBool(b.Options["sealed"], false)

	bb := deinterfacer.ToSliceOfStringToInterfaceMap(b.Options["buttons"])

	for b := range bb {
		bb[b]["enabled"] = deinterfacer.ToBool(bb[b]["enabled"], true)
		bb[b]["variant"] = deinterfacer.ToString(bb[b]["variant"], "primary")
		bb[b]["resourceType"] = deinterfacer.ToString(bb[b]["resourceType"], "compose:record")
	}

	b.Options["buttons"] = bb

	return b
}

// Get existing pages
func (pImp *Page) Get(handle string) (*types.Page, error) {
	handle = importer.NormalizeHandle(handle)
	if !importer.IsValidHandle(handle) {
		return nil, errors.New("invalid page handle")
	}

	return pImp.set.FindByHandle(handle), nil
}

func (pImp *Page) Store(ctx context.Context, k pageKeeper) (err error) {
	if err = pImp.storeChildren(ctx, "", k); err != nil {
		return
	}

	// We do that at the end - and save all pages with resolved references
	//
	// Many because internal page referencing from page blocks
	var refs uint
	for _, page := range pImp.set {
		if refs, err = pImp.resolveRefs(page); err != nil {
			return
		} else if refs > 0 {
			// make sure we do not get stale-data error
			page.UpdatedAt = nil
			if _, err = k.Update(page); err != nil {
				return errors.Wrap(err, "could not update resolved refs")
			}
		}
	}

	return
}

func (pImp *Page) storeChildren(ctx context.Context, parent string, k pageKeeper) (err error) {
	children, ok := pImp.tree[parent]
	if !ok {
		// No children...
		return nil
	}

	var parentPage *types.Page
	if parent != "" {
		parentPage, err = pImp.Get(parent)
		if err != nil {
			return
		} else if parentPage == nil {
			return errors.Errorf("could not load parent %q", parent)
		}
	}

	var page *types.Page

	for w, child := range children {
		if page, err = pImp.Get(child); err != nil {
			return
		}

		if parentPage != nil {
			page.SelfID = parentPage.ID
		}

		page.Weight = w

		if page.ID == 0 {
			page.NamespaceID = pImp.namespace.ID
			page, err = k.Create(page)
		} else if pImp.dirty[page.ID] {
			page, err = k.Update(page)
		}

		if err != nil {
			return
		}

		pImp.dirty[page.ID] = false
		if page.Handle == "" {
			continue
		}

		pImp.imp.permissions.UpdateResources(types.PagePermissionResource.String(), page.Handle, page.ID)

		if err = pImp.storeChildren(ctx, page.Handle, k); err != nil {
			return err
		}
	}

	return
}

// Resolve all refs for this page (page module, inside block)
//
// It counts number of resolved refs so that caller can know
// if there is anything to save
func (pImp *Page) resolveRefs(page *types.Page) (uint, error) {
	var refs uint

	var resolveModule = func(v interface{}) (string, error) {
		if m, err := pImp.getModule(deinterfacer.ToString(v)); err != nil {
			return "", err
		} else if m == nil {
			return "", fmt.Errorf("no such module (%v)", v)
		} else {
			refs++
			return strconv.FormatUint(m.ID, 10), nil
		}
	}

	var resolveModuleInMap = func(m map[string]interface{}) error {
		if h, ok := m["module"]; ok {
			if id, err := resolveModule(h); err != nil {
				return err
			} else {
				m["moduleID"] = id
				delete(m, "module")
			}
		}

		return nil
	}

	return refs, func() error {
		if moduleHandle, ok := pImp.modules[page.Handle]; ok {
			if refm, err := pImp.getModule(moduleHandle); err != nil || refm == nil {
				return errors.Errorf("could not load module %q for page %q (err: %v)",
					moduleHandle, page.Handle, err)
			} else {
				page.ModuleID = refm.ID
				refs++
			}
		}

		for i, b := range page.Blocks {
			if b.Options == nil {
				continue
			}

			if err := resolveModuleInMap(b.Options); err != nil {
				return fmt.Errorf("could not load module for page %q block #%d: %w", page.Handle, i+1, err)
			}

			if h, ok := b.Options["page"]; ok {
				if refp, err := pImp.Get(deinterfacer.ToString(h)); err != nil || refp == nil {
					return errors.Errorf("could not load page %q for page %q block #%d (err: %v)",
						h, page.Handle, i+1, err)
				} else {
					b.Options["pageID"] = strconv.FormatUint(refp.ID, 10)
					delete(b.Options, "page")
					refs++
				}
			}

			if h, ok := b.Options["chart"]; ok {
				if refc, err := pImp.getChart(deinterfacer.ToString(h)); err != nil || refc == nil {
					return errors.Errorf("could not load chart %q for page %q block #%d (err: %v)",
						h, page.Handle, i+1, err)
				} else {
					b.Options["chartID"] = strconv.FormatUint(refc.ID, 10)
					delete(b.Options, "chart")
					refs++
				}
			}

			switch b.Kind {
			case "Calendar":
				ff := make([]interface{}, 0)
				err := deinterfacer.Each(b.Options["feeds"], func(_ int, _ string, def interface{}) (err error) {
					feed := map[string]interface{}{}

					err = deinterfacer.Each(def, func(_ int, k string, v interface{}) error {
						switch k {
						case "options":
							if opt, ok := v.(map[string]interface{}); ok {
								if err := resolveModuleInMap(opt); err != nil {
									return fmt.Errorf("could not load resolve module for page %q Calendar block #%d: %w", page.Handle, i+1, err)
								}

								feed["options"] = opt
							}

						default:
							feed[k] = v
						}

						return nil
					})

					if err != nil {
						return err
					}

					ff = append(ff, feed)
					return nil
				})

				b.Options["feeds"] = ff

				if err != nil {
					return err
				}
			case "Metric":
				mm := make([]interface{}, 0)
				err := deinterfacer.Each(b.Options["metrics"], func(_ int, _ string, def interface{}) (err error) {
					metric := map[string]interface{}{}

					err = deinterfacer.Each(def, func(_ int, k string, v interface{}) error {
						switch k {
						case "module":
							if moduleID, err := resolveModule(v); err != nil {
								return fmt.Errorf("could not load resolve module for page %q Metric block #%d: %w", page.Handle, i+1, err)
							} else {
								metric["moduleID"] = moduleID
							}
						default:
							metric[k] = v
						}

						return nil
					})

					if err != nil {
						return err
					}

					mm = append(mm, metric)
					return nil
				})

				b.Options["metrics"] = mm

				if err != nil {
					return err
				}
			}
		}

		return nil
	}()
}
