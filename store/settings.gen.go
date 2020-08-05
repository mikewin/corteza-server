package store

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// store/settings.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/scenario"
	"github.com/cortezaproject/corteza-server/pkg/settings"
)

type (
	Settings interface {
		SearchSettings(ctx context.Context, f settings.Filter) (settings.ValueSet, settings.Filter, error)
		LookupSettingByNameOwnedBy(ctx context.Context, name string, owned_by uint64) (*settings.Value, error)
		CreateSetting(ctx context.Context, rr ...*settings.Value) error
		UpdateSetting(ctx context.Context, rr ...*settings.Value) error
		PartialUpdateSetting(ctx context.Context, onlyColumns []string, rr ...*settings.Value) error
		RemoveSetting(ctx context.Context, rr ...*settings.Value) error
		RemoveSettingByNameOwnedBy(ctx context.Context, name string, ownedBy uint64) error

		TruncateSettings(ctx context.Context) error
	}

	SettingsProvisioned interface {
		Settings
		ProvisionSettings() scenario.Executor
	}
)
