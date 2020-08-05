package tests

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
		var s = all.(store.ApplicationsUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeApplications()))
			req.NoError(s.TruncateApplications(ctx))
		})

		testApplications(t, s)
	})

	// Run generated tests for ComposeCharts
	t.Run("ComposeCharts", func(t *testing.T) {
		var s = all.(store.ComposeChartsUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeComposeCharts()))
			req.NoError(s.TruncateComposeCharts(ctx))
		})

		testComposeCharts(t, s)
	})

	// Run generated tests for ComposeModuleFields
	t.Run("ComposeModuleFields", func(t *testing.T) {
		var s = all.(store.ComposeModuleFieldsUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeComposeModuleFields()))
			req.NoError(s.TruncateComposeModuleFields(ctx))
		})

		testComposeModuleFields(t, s)
	})

	// Run generated tests for ComposeModules
	t.Run("ComposeModules", func(t *testing.T) {
		var s = all.(store.ComposeModulesUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeComposeModules()))
			req.NoError(s.TruncateComposeModules(ctx))
		})

		testComposeModules(t, s)
	})

	// Run generated tests for ComposeNamespaces
	t.Run("ComposeNamespaces", func(t *testing.T) {
		var s = all.(store.ComposeNamespacesUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeComposeNamespaces()))
			req.NoError(s.TruncateComposeNamespaces(ctx))
		})

		testComposeNamespaces(t, s)
	})

	// Run generated tests for ComposePages
	t.Run("ComposePages", func(t *testing.T) {
		var s = all.(store.ComposePagesUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeComposePages()))
			req.NoError(s.TruncateComposePages(ctx))
		})

		testComposePages(t, s)
	})

	// Run generated tests for Credentials
	t.Run("Credentials", func(t *testing.T) {
		var s = all.(store.CredentialsUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeCredentials()))
			req.NoError(s.TruncateCredentials(ctx))
		})

		testCredentials(t, s)
	})

	// Run generated tests for Reminders
	t.Run("Reminders", func(t *testing.T) {
		var s = all.(store.RemindersUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeReminders()))
			req.NoError(s.TruncateReminders(ctx))
		})

		testReminders(t, s)
	})

	// Run generated tests for Roles
	t.Run("Roles", func(t *testing.T) {
		var s = all.(store.RolesUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeRoles()))
			req.NoError(s.TruncateRoles(ctx))
		})

		testRoles(t, s)
	})

	// Run generated tests for Settings
	t.Run("Settings", func(t *testing.T) {
		var s = all.(store.SettingsUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeSettings()))
			req.NoError(s.TruncateSettings(ctx))
		})

		testSettings(t, s)
	})

	// Run generated tests for Attachment
	t.Run("Attachment", func(t *testing.T) {
		var s = all.(store.AttachmentsUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeAttachments()))
			req.NoError(s.TruncateAttachments(ctx))
		})

		testAttachment(t, s)
	})

	// Run generated tests for Users
	t.Run("Users", func(t *testing.T) {
		var s = all.(store.UsersUpgrader)

		t.Run("Upgrade", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(scenario.NewScenario(nil).Play(s.UpgradeUsers()))
			req.NoError(s.TruncateUsers(ctx))
		})

		testUsers(t, s)
	})

}
