package pgsql

import . "github.com/cortezaproject/corteza-server/pkg/scenario"

func (s *Store) UpgradeMessaging() Executor {
	return Do(
		s.UpgradeMessagingChannel(),
		s.UpgradeMessagingMessage(),
		s.UpgradeMessagingPermissionRules(),
		s.UpgradeMessagingSettings(),
	)
}

func (s *Store) UpgradeMessagingChannel() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.MessagingChannelTableDef()),
		sp.createTable(s.MessagingChannelMemberTableDef()),
	)
}

func (s *Store) UpgradeMessagingMessage() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.MessagingMessageTableDef()),
		sp.createTable(s.MessagingMessageAttachmentTableDef()),
		sp.createTable(s.MessagingMessageFlagTableDef()),
		sp.createTable(s.MessagingUnreadTableDef()),
		sp.createTable(s.MessagingAttachmentTableDef()),
	)
}

func (s *Store) UpgradeMessagingPermissionRules() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.MessagingSettingsTableDef()),
	)
}

func (s *Store) UpgradeMessagingSettings() Executor {
	sp := &storeUpgrade{s}

	return Do(
		sp.createTable(s.MessagingSettingsTableDef()),
	)
}
