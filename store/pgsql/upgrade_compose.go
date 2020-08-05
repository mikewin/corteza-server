package pgsql

import (
	. "github.com/cortezaproject/corteza-server/pkg/scenario"
)

func (s *Store) UpgradeCompose() Executor {
	return Do(
		s.UpgradeComposeNamespaces(),
		s.UpgradeComposePages(),
		s.UpgradeComposeCharts(),
		s.UpgradeComposeModules(),
		s.UpgradeComposeRecords(),
		s.UpgradeComposePermissionRules(),
		s.UpgradeComposeSettings(),
		s.UpgradeComposeAttachment(),
	)
}

func (s *Store) UpgradeComposeNamespaces() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.ComposeNamespaceTableDef()),
	)
}

func (s *Store) UpgradeComposePages() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.ComposePageTableDef()),
	)
}

func (s *Store) UpgradeComposeCharts() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.ComposeChartTableDef()),
	)
}

func (s *Store) UpgradeComposeModules() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.ComposeModuleTableDef()),
	)
}

func (s *Store) UpgradeComposeModuleFields() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.ComposeModuleFieldTableDef()),
	)
}

func (s *Store) UpgradeComposeRecords() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.ComposeRecordTableDef()),
		sp.createTable(s.ComposeRecordValueTableDef()),
	)
}

func (s *Store) UpgradeComposePermissionRules() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.ComposeSettingsTableDef()),
	)
}

func (s *Store) UpgradeComposeSettings() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.ComposeSettingsTableDef()),
	)
}

func (s *Store) UpgradeComposeAttachment() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.ComposeAttachmentTableDef()),
	)
}
