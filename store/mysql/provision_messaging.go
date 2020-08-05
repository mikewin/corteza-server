package mysql

import (
	. "github.com/cortezaproject/corteza-server/pkg/scenario"
)

func (s *Store) ProvisionMessaging() Executor {
	return Do(
		s.ProvisionMessagingChannel(),
		s.ProvisionMessagingMessage(),
		s.ProvisionMessagingPermissionRules(),
		s.ProvisionMessagingSettings(),
	)
}

func (s *Store) ProvisionMessagingChannel() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.MessagingChannelTableDef()),
		sp.createTable(s.MessagingChannelMemberTableDef()),
	)
}

func (s *Store) ProvisionMessagingMessage() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.MessagingMessageTableDef()),
		sp.createTable(s.MessagingMessageAttachmentTableDef()),
		sp.createTable(s.MessagingMessageFlagTableDef()),
		sp.createTable(s.MessagingUnreadTableDef()),
		sp.createTable(s.MessagingAttachmentTableDef()),
	)
}

func (s *Store) ProvisionMessagingPermissionRules() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.MessagingSettingsTableDef()),
	)
}

func (s *Store) ProvisionMessagingSettings() Executor {
	sp := &storeProvision{s}

	return Do(
		sp.createTable(s.MessagingSettingsTableDef()),
	)
}
