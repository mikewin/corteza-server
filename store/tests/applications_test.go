package test_store

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/id"
	"github.com/cortezaproject/corteza-server/pkg/rh"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func testApplications(t *testing.T, s store.Applications) {
	var (
		ctx         = context.Background()
		req         = require.New(t)
		application *types.Application

		newApp = func(name string) *types.Application {
			// minimum data set for new application
			return &types.Application{
				ID:        id.Next(),
				CreatedAt: time.Now(),
				Name:      name,
				Unify:     &types.ApplicationUnify{},
			}
		}
	)

	t.Run("create", func(t *testing.T) {
		application = newApp("ApplicationCRUD")
		req.NoError(s.CreateApplication(ctx, application))
	})

	t.Run("lookup by ID", func(t *testing.T) {
		fetched, err := s.LookupApplicationByID(ctx, application.ID)
		req.NoError(err)
		req.Equal(application.Name, fetched.Name)
		req.Equal(application.ID, fetched.ID)
		req.NotNil(fetched.CreatedAt)
		req.Nil(fetched.UpdatedAt)
		req.Nil(fetched.DeletedAt)
	})

	t.Run("update", func(t *testing.T) {
		application = &types.Application{
			ID:        application.ID,
			CreatedAt: application.CreatedAt,
			Name:      "ApplicationCRUD+2",
			Unify:     application.Unify,
		}
		req.NoError(s.UpdateApplication(ctx, application))

		updated, err := s.LookupApplicationByID(ctx, application.ID)
		req.NoError(err)
		req.Equal(application.Name, updated.Name)
	})

	t.Run("search", func(t *testing.T) {
		prefill := []*types.Application{
			newApp("app-one-one"),
			newApp("app-one-two"),
			newApp("app-two-one"),
			newApp("app-two-two"),
			newApp("app-two-deleted"),
		}

		count := len(prefill)

		prefill[4].DeletedAt = &prefill[4].CreatedAt
		valid := count - 1

		req.NoError(s.TruncateApplications(ctx))
		req.NoError(s.CreateApplication(ctx, prefill...))

		// search for all valid
		set, f, err := s.SearchApplications(ctx, types.ApplicationFilter{})
		req.NoError(err)
		req.Len(set, valid) // we've deleted one
		req.Equal(valid, int(f.Count))

		// search for ALL
		set, f, err = s.SearchApplications(ctx, types.ApplicationFilter{Deleted: rh.FilterStateInclusive})
		req.NoError(err)
		req.Len(set, count) // we've deleted one

		// search for deleted only
		set, f, err = s.SearchApplications(ctx, types.ApplicationFilter{Deleted: rh.FilterStateExclusive})
		req.NoError(err)
		req.Len(set, 1) // we've deleted one

		set, f, err = s.SearchApplications(ctx, types.ApplicationFilter{Name: "app-two-one"})
		req.NoError(err)
		req.Len(set, 1)

		// find all prefixed
		set, f, err = s.SearchApplications(ctx, types.ApplicationFilter{Query: "app-two"})
		req.NoError(err)
		req.Len(set, 2)
	})
}
