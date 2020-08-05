package mysql

import (
	. "github.com/cortezaproject/corteza-server/pkg/scenario"
)

func (s *Store) ProvisionCompose() Executor {
	return Do(
		s.ProvisionComposeNamespaces(),
		s.ProvisionComposePages(),
		s.ProvisionComposeCharts(),
		s.ProvisionComposeModules(),
		s.ProvisionComposeRecords(),
		s.ProvisionComposePermissionRules(),
		s.ProvisionComposeSettings(),
		s.ProvisionComposeAttachment(),
	)
}

func (s *Store) ProvisionComposeNamespaces() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.ComposeNamespaceTableDef()),
	)
}

func (s *Store) ProvisionComposePages() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.ComposePageTableDef()),
	)
}

func (s *Store) ProvisionComposeCharts() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.ComposeChartTableDef()),
	)
}

func (s *Store) ProvisionComposeModules() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.ComposeModuleTableDef()),
	)
}

func (s *Store) ProvisionComposeModuleFields() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.ComposeModuleFieldTableDef()),
	)
}

func (s *Store) ProvisionComposeRecords() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.ComposeRecordTableDef()),
		sp.createTable(s.ComposeRecordValueTableDef()),
	)
}

func (s *Store) ProvisionComposePermissionRules() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.ComposeSettingsTableDef()),
	)
}

func (s *Store) ProvisionComposeSettings() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.ComposeSettingsTableDef()),
	)
}

func (s *Store) ProvisionComposeAttachment() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.ComposeAttachmentTableDef()),
	)
}
