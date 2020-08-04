package bulk

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
// Definitions file that controls how this file is generated:
// store/settings.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/settings"
	"github.com/cortezaproject/corteza-server/store"
)

type (
	settingCreate struct {
		Done chan struct{}
		res  *settings.Value
		err  error
	}

	settingUpdate struct {
		Done chan struct{}
		res  *settings.Value
		err  error
	}

	settingRemove struct {
		Done chan struct{}
		res  *settings.Value
		err  error
	}
)

// CreateSetting creates a new Setting
// create job that can be pushed to store's transaction handler
func CreateSetting(res *settings.Value) *settingCreate {
	return &settingCreate{res: res}
}

// Do Executes settingCreate job
func (j *settingCreate) Do(ctx context.Context, s store.Interface) error {
	j.err = s.CreateSetting(ctx, j.res)
	j.Done <- struct{}{}
	return j.err
}

// UpdateSetting creates a new Setting
// update job that can be pushed to store's transaction handler
func UpdateSetting(res *settings.Value) *settingUpdate {
	return &settingUpdate{res: res}
}

// Do Executes settingUpdate job
func (j *settingUpdate) Do(ctx context.Context, s store.Interface) error {
	j.err = s.UpdateSetting(ctx, j.res)
	j.Done <- struct{}{}
	return j.err
}

// RemoveSetting creates a new Setting
// remove job that can be pushed to store's transaction handler
func RemoveSetting(res *settings.Value) *settingRemove {
	return &settingRemove{res: res}
}

// Do Executes settingRemove job
func (j *settingRemove) Do(ctx context.Context, s store.Interface) error {
	j.err = s.RemoveSetting(ctx, j.res)
	j.Done <- struct{}{}
	return j.err
}
