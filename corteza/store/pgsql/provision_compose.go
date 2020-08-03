package pgsql

import "github.com/cortezaproject/corteza-server/corteza/store/provisioner"

func (s *Store) ProvisionCompose() provisioner.Executor {
	return provisioner.Do(
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

func (s *Store) ProvisionComposeNamespaces() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.ComposeNamespaceTableDef()),
	)
}

func (s *Store) ProvisionComposePages() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.ComposePageTableDef()),
	)
}

func (s *Store) ProvisionComposeCharts() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.ComposeChartTableDef()),
	)
}

func (s *Store) ProvisionComposeModules() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.ComposeModuleTableDef()),
	)
}

func (s *Store) ProvisionComposeModuleFields() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.ComposeModuleFieldTableDef()),
	)
}

func (s *Store) ProvisionComposeRecords() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.ComposeRecordTableDef()),
		sp.createTable(s.ComposeRecordValueTableDef()),
	)
}

func (s *Store) ProvisionComposePermissionRules() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.ComposeSettingsTableDef()),
	)
}

func (s *Store) ProvisionComposeSettings() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.ComposeSettingsTableDef()),
	)
}

func (s *Store) ProvisionComposeAttachment() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.ComposeAttachmentTableDef()),
	)
}
