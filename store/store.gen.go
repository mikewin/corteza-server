package store

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

type (
	// Interface combines interfaces of all supported store interfaces
	Interface interface {
		Applications
		ComposeCharts
		ComposeModuleFields
		ComposeModules
		ComposeNamespaces
		ComposePages
		Credentials
		Reminders
		Roles
		Settings
		Attachments
		Users
	}

	// Interface combines interfaces of all supported store interfaces
	Upgrader interface {
		ApplicationsUpgrader
		ComposeChartsUpgrader
		ComposeModuleFieldsUpgrader
		ComposeModulesUpgrader
		ComposeNamespacesUpgrader
		ComposePagesUpgrader
		CredentialsUpgrader
		RemindersUpgrader
		RolesUpgrader
		SettingsUpgrader
		AttachmentsUpgrader
		UsersUpgrader
	}
)
