package mysql

import (
	. "github.com/cortezaproject/corteza-server/pkg/scenario"
)

func (s *Store) ProvisionCore() Executor {
	return Do(
		s.ProvisionUsers(),
		s.ProvisionActionLog(),
		s.ProvisionPermissionRules(),
		s.ProvisionSettings(),
		s.ProvisionAttachments(),
		s.ProvisionApplications(),
		s.ProvisionReminders(),
	)
}

func (s *Store) ProvisionActionLog() Executor {
	return (&storeProvision{s}).createTable(
		s.ActionLogTableDef(),
	)
}

func (s *Store) ProvisionPermissionRules() Executor {
	return (&storeProvision{s}).createTable(
		s.PermissionRulesTableDef(),
	)
}

func (s *Store) ProvisionSettings() Executor {
	return (&storeProvision{s}).createTable(
		s.SettingsTableDef(),
	)
}

func (s *Store) ProvisionAttachments() Executor {
	return (&storeProvision{s}).createTable(
		s.AttachmentsTableDef(),
	)
}

func (s *Store) ProvisionUsers() Executor {
	return (&storeProvision{s}).createTable(
		s.UsersTableDef(),
	//dropColumn(sysUser, "rel_organisation"),
	//dropColumn(sysUser, "rel_user_id"),
	//addColumn(sysRole, "created_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
	//addColumn(sysRole, "updated_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
	//addColumn(sysRole, "deleted_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
	)
}

func (s *Store) ProvisionCredentials() Executor {
	return (&storeProvision{s}).createTable(
		s.CredentialsTableDef(),
	)
}

func (s *Store) ProvisionRoles() Executor {
	return (&storeProvision{s}).createTable(
		s.RolesTableDef(),
	)
}

func (s *Store) ProvisionApplications() Executor {
	return (&storeProvision{s}).createTable(
		s.ApplicationsTableDef(),
	)
}

func (s *Store) ProvisionReminders() Executor {
	return (&storeProvision{s}).createTable(
		s.RemindersTableDef(),
	)
}
