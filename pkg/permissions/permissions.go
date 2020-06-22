package permissions

// General permission stuff, types, constants

type (
	Operation string
	Access    int

	// CheckAccessFunc function.
	CheckAccessFunc func() Access

	Whitelist struct {
		// Index is used for fast lookups
		index map[Resource]map[Operation]bool

		// we need this to maintain a stable order of res/ops
		rules RuleSet
	}

	whitelistFlatten struct {
		Resource  `json:"resource"`
		Operation `json:"operation"`
	}
)

const (
	// Hardcoded Role ID for everyone
	EveryoneRoleID uint64 = 1

	// Hardcoded ID for Admin role
	AdminsRoleID uint64 = 2

	// OwnersDynamicRoleID role is dynamically assigned when current user is owner of the resource
	OwnersDynamicRoleID uint64 = 10000

	// CreatorsDynamicRoleID role is dynamically assigned when current user created the resource
	CreatorsDynamicRoleID uint64 = 10010

	// UpdatorsDynamicRoleID role is dynamically assigned when current user updated the resource
	UpdatersDynamicRoleID uint64 = 10011

	// DeletersDynamicRoleID role is dynamically assigned when current user deleted the resource
	DeletersDynamicRoleID uint64 = 10012

	// MembersDynamicRoleID role is dynamically assigned when current user member of the resource
	MembersDynamicRoleID uint64 = 10020
)

func (op Operation) String() string {
	return string(op)
}

func (a Access) String() string {
	switch a {
	case Allow:
		return "allow"
	case Deny:
		return "deny"
	default:
		return "inherit"
	}
}

// Bool convers boolean true to Allow and false to Deny
func BoolToCheckFunc(isTrue bool) CheckAccessFunc {
	return func() Access {
		if isTrue {
			return Allow
		}

		return Deny
	}
}

func (a *Access) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "allow":
		*a = Allow
	case "deny":
		*a = Deny
	default:
		*a = Inherit
	}
	return nil
}

func (a Access) MarshalJSON() ([]byte, error) {
	return []byte(`"` + a.String() + `"`), nil
}

func Allowed() Access {
	return Allow
}

func Denied() Access {
	return Deny
}

func (wl *Whitelist) Set(r Resource, oo ...Operation) {
	if wl.index == nil {
		wl.index = map[Resource]map[Operation]bool{}
	}

	wl.index[r] = map[Operation]bool{}

	for _, o := range oo {
		wl.index[r][o] = true
		wl.rules = append(wl.rules, &Rule{Resource: r, Operation: o})
	}
}

func (wl Whitelist) Check(rule *Rule) bool {
	if rule == nil {
		return false
	}

	res := rule.Resource.TrimID()

	if _, ok := wl.index[res]; !ok {
		return false
	}

	return wl.index[res][rule.Operation]
}

// Flatten casts list of operations for each resource from map to slice and creates more output friendly format
func (wl Whitelist) Flatten() []whitelistFlatten {
	var (
		wlf = []whitelistFlatten{}
	)
	for _, r := range wl.rules {
		wlf = append(wlf, whitelistFlatten{r.Resource, r.Operation})
	}

	return wlf
}

// DynamicRoles is a utility function that compares
// given u with each odd element in cc
// and returns even element on a match
func DynamicRoles(u uint64, cc ...uint64) (rr []uint64) {
	var l = len(cc)

	if l%2 == 1 {
		panic("expecting even number of id/dynamic-role pairs")
	}

	rr = make([]uint64, 0, l/2)

	for i := 0; i < l; i += 2 {
		if cc[i] == u {
			rr = append(rr, cc[i+1])
		}
	}

	return
}
