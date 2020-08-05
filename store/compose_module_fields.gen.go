package store

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// store/compose_module_fields.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/scenario"
)

type (
	ComposeModuleFields interface {
		CreateComposeModuleField(ctx context.Context, rr ...*types.ModuleField) error
		UpdateComposeModuleField(ctx context.Context, rr ...*types.ModuleField) error
		PartialUpdateComposeModuleField(ctx context.Context, onlyColumns []string, rr ...*types.ModuleField) error
		RemoveComposeModuleField(ctx context.Context, rr ...*types.ModuleField) error
		RemoveComposeModuleFieldByName(ctx context.Context, name string) error

		TruncateComposeModuleFields(ctx context.Context) error
	}

	ComposeModuleFieldsUpgrader interface {
		ComposeModuleFields
		UpgradeComposeModuleFields() scenario.Executor
	}
)
