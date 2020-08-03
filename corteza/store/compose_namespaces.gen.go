package store

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// corteza/store/compose_namespaces.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/corteza/store/provisioner"
)

type (
	ComposeNamespaces interface {
		SearchComposeNamespaces(ctx context.Context, f types.NamespaceFilter) (types.NamespaceSet, types.NamespaceFilter, error)
		LookupComposeNamespaceBySlug(ctx context.Context, slug string) (*types.Namespace, error)
		CreateComposeNamespace(ctx context.Context, rr ...*types.Namespace) error
		UpdateComposeNamespace(ctx context.Context, rr ...*types.Namespace) error
		PartialUpdateComposeNamespace(ctx context.Context, onlyColumns []string, rr ...*types.Namespace) error
		RemoveComposeNamespace(ctx context.Context, rr ...*types.Namespace) error
		RemoveComposeNamespaceByID(ctx context.Context, ID uint64) error

		TruncateComposeNamespaces(ctx context.Context) error
	}

	ComposeNamespacesProvisioned interface {
		ComposeNamespaces
		ProvisionComposeNamespaces() provisioner.Executor
	}
)
