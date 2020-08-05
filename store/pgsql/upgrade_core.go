package pgsql

import (
	. "github.com/cortezaproject/corteza-server/pkg/scenario"
)

func (s *Store) UpgradeCore() Executor {
	return Do(
		s.UpgradeUsers(),
		s.UpgradeActionLog(),
		s.UpgradePermissionRules(),
		s.UpgradeSettings(),
		s.UpgradeAttachments(),
		s.UpgradeApplications(),
		s.UpgradeReminders(),
	)
}

func (s *Store) UpgradeActionLog() Executor {
	return (&storeUpgrade{s}).createTable(
		s.ActionLogTableDef(),
	)
}

func (s *Store) UpgradePermissionRules() Executor {
	return (&storeUpgrade{s}).createTable(
		s.PermissionRulesTableDef(),
	)
}

func (s *Store) UpgradeSettings() Executor {
	return (&storeUpgrade{s}).createTable(
		s.SettingsTableDef(),
	)
}

func (s *Store) UpgradeAttachments() Executor {
	return (&storeUpgrade{s}).createTable(
		s.AttachmentsTableDef(),
	)
}

func (s *Store) UpgradeUsers() Executor {
	return (&storeUpgrade{s}).createTable(
		s.UsersTableDef(),
	//dropColumn(sysUser, "rel_organisation"),
	//dropColumn(sysUser, "rel_user_id"),
	//addColumn(sysRole, "created_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
	//addColumn(sysRole, "updated_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
	//addColumn(sysRole, "deleted_by", "BIGINT UNSIGNED NOT NULL DEFAULT 0"),
	)
}

func (s *Store) UpgradeCredentials() Executor {
	return (&storeUpgrade{s}).createTable(
		s.CredentialsTableDef(),
	)
}

func (s *Store) UpgradeRoles() Executor {
	return (&storeUpgrade{s}).createTable(
		s.RolesTableDef(),
	)
}

func (s *Store) UpgradeApplications() Executor {
	return (&storeUpgrade{s}).createTable(
		s.ApplicationsTableDef(),
	)
}

func (s *Store) UpgradeReminders() Executor {
	return (&storeUpgrade{s}).createTable(
		s.RemindersTableDef(),
	)
}
