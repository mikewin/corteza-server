package test_store

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/scenario"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/stretchr/testify/require"
	"testing"
)

func testAllGenerated(t *testing.T, all interface{}) {

	// Run generated tests for Applications
	t.Run("Applications", func(t *testing.T) {
		var s = all.(store.ApplicationsProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionApplications()))
			req.NoError(s.TruncateApplications(ctx))
		})

		testApplications(t, s)
	})

	// Run generated tests for ComposeCharts
	t.Run("ComposeCharts", func(t *testing.T) {
		var s = all.(store.ComposeChartsProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionComposeCharts()))
			req.NoError(s.TruncateComposeCharts(ctx))
		})

		testComposeCharts(t, s)
	})

	// Run generated tests for ComposeModuleFields
	t.Run("ComposeModuleFields", func(t *testing.T) {
		var s = all.(store.ComposeModuleFieldsProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionComposeModuleFields()))
			req.NoError(s.TruncateComposeModuleFields(ctx))
		})

		testComposeModuleFields(t, s)
	})

	// Run generated tests for ComposeModules
	t.Run("ComposeModules", func(t *testing.T) {
		var s = all.(store.ComposeModulesProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionComposeModules()))
			req.NoError(s.TruncateComposeModules(ctx))
		})

		testComposeModules(t, s)
	})

	// Run generated tests for ComposeNamespaces
	t.Run("ComposeNamespaces", func(t *testing.T) {
		var s = all.(store.ComposeNamespacesProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionComposeNamespaces()))
			req.NoError(s.TruncateComposeNamespaces(ctx))
		})

		testComposeNamespaces(t, s)
	})

	// Run generated tests for ComposePages
	t.Run("ComposePages", func(t *testing.T) {
		var s = all.(store.ComposePagesProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionComposePages()))
			req.NoError(s.TruncateComposePages(ctx))
		})

		testComposePages(t, s)
	})

	// Run generated tests for Credentials
	t.Run("Credentials", func(t *testing.T) {
		var s = all.(store.CredentialsProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionCredentials()))
			req.NoError(s.TruncateCredentials(ctx))
		})

		testCredentials(t, s)
	})

	// Run generated tests for Reminders
	t.Run("Reminders", func(t *testing.T) {
		var s = all.(store.RemindersProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionReminders()))
			req.NoError(s.TruncateReminders(ctx))
		})

		testReminders(t, s)
	})

	// Run generated tests for Roles
	t.Run("Roles", func(t *testing.T) {
		var s = all.(store.RolesProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionRoles()))
			req.NoError(s.TruncateRoles(ctx))
		})

		testRoles(t, s)
	})

	// Run generated tests for Settings
	t.Run("Settings", func(t *testing.T) {
		var s = all.(store.SettingsProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionSettings()))
			req.NoError(s.TruncateSettings(ctx))
		})

		testSettings(t, s)
	})

	// Run generated tests for Attachment
	t.Run("Attachment", func(t *testing.T) {
		var s = all.(store.AttachmentsProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionAttachments()))
			req.NoError(s.TruncateAttachments(ctx))
		})

		testAttachment(t, s)
	})

	// Run generated tests for Users
	t.Run("Users", func(t *testing.T) {
		var s = all.(store.UsersProvisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(runner.NewScenario(nil).Run(s.ProvisionUsers()))
			req.NoError(s.TruncateUsers(ctx))
		})

		testUsers(t, s)
	})

}
