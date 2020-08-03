package pgsql

import (
	"github.com/cortezaproject/corteza-server/corteza/store/provisioner"
)

func (s *Store) ProvisionCore() provisioner.Executor {
	return provisioner.Do(
		s.ProvisionUsers(),
		s.ProvisionActionLog(),
		s.ProvisionPermissionRules(),
		s.ProvisionSettings(),
		s.ProvisionAttachments(),
		s.ProvisionApplications(),
		s.ProvisionReminders(),
	)
}

func (s *Store) ProvisionActionLog() provisioner.Executor {
	return (&storeProvision{s}).createTable(
		s.ActionLogTableDef(),
	)
}

func (s *Store) ProvisionPermissionRules() provisioner.Executor {
	return (&storeProvision{s}).createTable(
		s.PermissionRulesTableDef(),
	)
}

func (s *Store) ProvisionSettings() provisioner.Executor {
	return (&storeProvision{s}).createTable(
		s.SettingsTableDef(),
	)
}

func (s *Store) ProvisionAttachments() provisioner.Executor {
	return (&storeProvision{s}).createTable(
		s.AttachmentsTableDef(),
	)
}

func (s *Store) ProvisionUsers() provisioner.Executor {
	return (&storeProvision{s}).createTable(
		s.UsersTableDef(),
	//dropColumn(sysUser, "rel_organisation"),
	//dropColumn(sysUser, "rel_user_id"),
	//addColumn(sysRole, "created_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
	//addColumn(sysRole, "updated_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
	//addColumn(sysRole, "deleted_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
	)
}

func (s *Store) ProvisionCredentials() provisioner.Executor {
	return (&storeProvision{s}).createTable(
		s.CredentialsTableDef(),
	)
}

func (s *Store) ProvisionRoles() provisioner.Executor {
	return (&storeProvision{s}).createTable(
		s.RolesTableDef(),
	)
}

func (s *Store) ProvisionApplications() provisioner.Executor {
	return (&storeProvision{s}).createTable(
		s.ApplicationsTableDef(),
	)
}

func (s *Store) ProvisionReminders() provisioner.Executor {
	return (&storeProvision{s}).createTable(
		s.RemindersTableDef(),
	)
}
