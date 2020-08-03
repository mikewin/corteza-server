package mysql

import "github.com/cortezaproject/corteza-server/corteza/store/provisioner"

func (s *Store) ProvisionMessaging() provisioner.Executor {
	return provisioner.Do(
		s.ProvisionMessagingChannel(),
		s.ProvisionMessagingMessage(),
		s.ProvisionMessagingPermissionRules(),
		s.ProvisionMessagingSettings(),
	)
}

func (s *Store) ProvisionMessagingChannel() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.MessagingChannelTableDef()),
		sp.createTable(s.MessagingChannelMemberTableDef()),
	)
}

func (s *Store) ProvisionMessagingMessage() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.MessagingMessageTableDef()),
		sp.createTable(s.MessagingMessageAttachmentTableDef()),
		sp.createTable(s.MessagingMessageFlagTableDef()),
		sp.createTable(s.MessagingUnreadTableDef()),
		sp.createTable(s.MessagingAttachmentTableDef()),
	)
}

func (s *Store) ProvisionMessagingPermissionRules() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.MessagingSettingsTableDef()),
	)
}

func (s *Store) ProvisionMessagingSettings() provisioner.Executor {
	sp := &storeProvision{s}

	return provisioner.Do(
		sp.createTable(s.MessagingSettingsTableDef()),
	)
}
