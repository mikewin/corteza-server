package rdbms

import (
	. "github.com/cortezaproject/corteza-server/store/rdbms/schema"
)

func (s *Store) UsersTableDef() *Table {
	return TableDef("sys_user",
		AddID(),
		ColumnDef("email", ColumnTypeText),
		ColumnDef("email_confirmed", ColumnTypeBoolean, DefaultValue("false")),
		ColumnDef("username", ColumnTypeText),
		ColumnDef("name", ColumnTypeText),
		ColumnDef("handle", ColumnTypeText),
		ColumnDef("kind", ColumnTypeText, ColumnTypeLength(8)),
		ColumnDef("meta", ColumnTypeJson),
		ColumnDef("suspended_at", ColumnTypeTimestamp, Null),
		CUDTimestamps,
	)
}
func (s *Store) CredentialsTableDef() *Table {
	return TableDef(`sys_credentials`,
		AddID(),
		ColumnDef("rel_owner", ColumnTypeIdentifier),
		ColumnDef("label", ColumnTypeText),
		ColumnDef("kind", ColumnTypeText),
		ColumnDef("credentials", ColumnTypeText),
		ColumnDef("meta", ColumnTypeJson),
		ColumnDef("expires_at", ColumnTypeTimestamp, Null),
		ColumnDef("last_used_at", ColumnTypeTimestamp, Null),
		CUDTimestamps,
	)
}
func (s *Store) RolesTableDef() *Table {
	return TableDef(`sys_role`,
		AddID(),
		ColumnDef("name", ColumnTypeText),
		ColumnDef("handle", ColumnTypeText),
		ColumnDef("archived_at", ColumnTypeTimestamp, Null),
		CUDTimestamps,
	)
}
func (s *Store) RoleMembersTableDef() *Table {
	return TableDef(`sys_role_member`,
		ColumnDef("rel_role", ColumnTypeIdentifier),
		ColumnDef("rel_user", ColumnTypeIdentifier),
	)
}

func (s *Store) ApplicationsTableDef() *Table {
	return TableDef("sys_application",
		AddID(),
		ColumnDef("name", ColumnTypeText),
		ColumnDef("enabled", ColumnTypeBoolean, DefaultValue("true")),
		ColumnDef("unify", ColumnTypeJson),
		ColumnDef("rel_owner", ColumnTypeIdentifier),
		CUDTimestamps,
	)
}

func (s *Store) RemindersTableDef() *Table {
	return TableDef("sys_reminder",
		AddID(),
		ColumnDef("resource", ColumnTypeText),
		ColumnDef("payload", ColumnTypeJson),
		ColumnDef("snooze_count", ColumnTypeInteger, DefaultValue("0")),
		ColumnDef("assigned_to", ColumnTypeIdentifier, DefaultValue("0")),
		ColumnDef("assigned_by", ColumnTypeIdentifier, DefaultValue("0")),
		ColumnDef("assigned_at", ColumnTypeTimestamp),
		ColumnDef("remind_at", ColumnTypeTimestamp, Null),
		ColumnDef("dismissed_at", ColumnTypeTimestamp, Null),
		ColumnDef("dismissed_by", ColumnTypeIdentifier, DefaultValue("0")),
		CUDTimestamps,
	)
}

func (s *Store) AttachmentsTableDef() *Table {
	return TableDef("sys_attachment",
		AddID(),
		ColumnDef("rel_owner", ColumnTypeIdentifier),
		ColumnDef("kind", ColumnTypeText),
		ColumnDef("url", ColumnTypeText),
		ColumnDef("preview_url", ColumnTypeText),
		ColumnDef("name", ColumnTypeText),
		ColumnDef("meta", ColumnTypeJson),
		CUDTimestamps,
	)
}

func (s *Store) ActionLogTableDef() *Table {
	return TableDef("sys_action_log",
		ColumnDef("ts", ColumnTypeTimestamp, DefaultValue("NOW()")),
		ColumnDef("actor_ip_addr", ColumnTypeText, ColumnTypeLength(15)),
		ColumnDef("actor_id", ColumnTypeIdentifier),
		ColumnDef("request_origin", ColumnTypeText),
		ColumnDef("request_id", ColumnTypeText),
		ColumnDef("resource", ColumnTypeText),
		ColumnDef("action", ColumnTypeText),
		ColumnDef("error", ColumnTypeText),
		ColumnDef("severity", ColumnTypeInteger),
		ColumnDef("description", ColumnTypeText),
		ColumnDef("meta", ColumnTypeJson),

		// @todo KEY ts (ts DESC),
		// @todo KEY request_origin (request_origin),
		// @todo KEY actor_id (actor_id),
		// @todo KEY resource (resource),
		// @todo KEY action (action)
	)
}

func (s *Store) PermissionRulesTableDef() *Table {
	return TableDef("sys_permission_rules",
		ColumnDef("rel_role", ColumnTypeIdentifier),
		ColumnDef("resource", ColumnTypeText),
		ColumnDef("operation", ColumnTypeText),
		ColumnDef("access", ColumnTypeInteger),

		SetPrimaryKey("rel_role", "resource", "operation", "access"),
	)
}

func (s *Store) SettingsTableDef() *Table {
	return TableDef("settings",
		ColumnDef("rel_owner", ColumnTypeIdentifier),
		ColumnDef("name", ColumnTypeText, ColumnTypeLength(200)),
		ColumnDef("value", ColumnTypeJson),
		ColumnDef("updated_by", ColumnTypeIdentifier),
		ColumnDef("updated_at", ColumnTypeTimestamp),

		SetPrimaryKey("name", "rel_owner"),
	)
}

func (s *Store) ComposeAttachmentTableDef() *Table {
	// @todo merge with general attachment table

	return TableDef("compose_attachment",
		AddID(),
		ColumnDef("rel_namespace", ColumnTypeIdentifier),
		ColumnDef("rel_owner", ColumnTypeIdentifier),
		ColumnDef("kind", ColumnTypeText),
		ColumnDef("url", ColumnTypeText),
		ColumnDef("preview_url", ColumnTypeText),
		ColumnDef("name", ColumnTypeText),
		ColumnDef("meta", ColumnTypeJson),

		CUDTimestamps,
		//  KEY rel_namespace (rel_namespace)
	)
}

func (s *Store) ComposeChartTableDef() *Table {
	return TableDef("compose_chart",
		AddID(),
		ColumnDef("handle", ColumnTypeText),
		ColumnDef("rel_namespace", ColumnTypeIdentifier),
		ColumnDef("name", ColumnTypeText),
		ColumnDef("config", ColumnTypeJson),
		CUDTimestamps,

		//  KEY rel_namespace (rel_namespace)
	)
}

func (s *Store) ComposeModuleTableDef() *Table {
	return TableDef("compose_module",
		AddID(),
		ColumnDef("handle", ColumnTypeText),
		ColumnDef("rel_namespace", ColumnTypeIdentifier),
		ColumnDef("name", ColumnTypeText),
		ColumnDef("json", ColumnTypeJson),
		CUDTimestamps,
	)
}

func (s *Store) ComposeModuleFieldTableDef() *Table {
	return TableDef("compose_module_field",
		AddID(),

		ColumnDef("rel_module", ColumnTypeIdentifier),
		ColumnDef("place", ColumnTypeInteger),
		ColumnDef("kind", ColumnTypeText),
		ColumnDef("options", ColumnTypeJson),
		ColumnDef("default_value", ColumnTypeJson),
		ColumnDef("name", ColumnTypeText, ColumnTypeLength(64)),
		ColumnDef("label", ColumnTypeText),
		ColumnDef("is_private", ColumnTypeBoolean),
		ColumnDef("is_required", ColumnTypeBoolean),
		ColumnDef("is_visible", ColumnTypeBoolean),
		ColumnDef("is_multi", ColumnTypeBoolean),

		CUDTimestamps,
	)
}

func (s *Store) ComposeNamespaceTableDef() *Table {
	return TableDef("compose_namespace",
		AddID(),
		ColumnDef("name", ColumnTypeText),
		ColumnDef("slug", ColumnTypeText),
		ColumnDef("enabled", ColumnTypeBoolean),
		ColumnDef("meta", ColumnTypeJson),
		CUDTimestamps,
	)
}

func (s *Store) ComposePageTableDef() *Table {
	return TableDef("compose_page",
		AddID(),
		ColumnDef("title", ColumnTypeText),
		ColumnDef("handle", ColumnTypeText),
		ColumnDef("description", ColumnTypeText),
		ColumnDef("rel_namespace", ColumnTypeIdentifier),
		ColumnDef("rel_module", ColumnTypeIdentifier),
		ColumnDef("self_id", ColumnTypeIdentifier),
		ColumnDef("blocks", ColumnTypeJson),
		ColumnDef("visible", ColumnTypeBoolean),
		ColumnDef("weight", ColumnTypeInteger),
		CUDTimestamps,
	//  KEY module_id (rel_module),
	//  KEY self_id (self_id),
	//  KEY rel_namespace (rel_namespace)
	)
}

func (s *Store) ComposePermissionRulesTableDef() *Table {
	return TableDef("compose_permission_rules",
		ColumnDef("rel_role", ColumnTypeIdentifier),
		ColumnDef("resource", ColumnTypeText),
		ColumnDef("operation", ColumnTypeText),
		ColumnDef("access", ColumnTypeInteger),

		SetPrimaryKey("rel_role", "resource", "operation", "access"),
	)
}

func (s *Store) ComposeRecordTableDef() *Table {
	return TableDef("compose_record",
		AddID(),
		ColumnDef("rel_namespace", ColumnTypeIdentifier),
		ColumnDef("rel_module", ColumnTypeIdentifier),
		ColumnDef("owned_by", ColumnTypeIdentifier),
		CUDTimestamps,
		CUDUsers,

		//  PRIMARY KEY (id),
		//  KEY user_id (owned_by),
		//  KEY rel_module (rel_module),
		//  KEY rel_namespace (rel_namespace)
	)
}

func (s *Store) ComposeRecordValueTableDef() *Table {
	return TableDef("compose_record_value",
		ColumnDef("record_id", ColumnTypeIdentifier),
		ColumnDef("name", ColumnTypeText, ColumnTypeLength(64)),
		ColumnDef("ref", ColumnTypeIdentifier),
		ColumnDef("place", ColumnTypeInteger),
		ColumnDef("deleted_at", ColumnTypeTimestamp, Null),

		SetPrimaryKey("record_id", "name", "place"),
		// KEY crm_record_value_ref (ref)
	)
}

func (s *Store) ComposeSettingsTableDef() *Table {
	// @todo merge with general settings table
	return TableDef("compose_settings",
		ColumnDef("rel_owner", ColumnTypeIdentifier),
		ColumnDef("name", ColumnTypeText, ColumnTypeLength(200)),
		ColumnDef("value", ColumnTypeJson),
		ColumnDef("updated_by", ColumnTypeIdentifier),
		ColumnDef("updated_at", ColumnTypeTimestamp),

		SetPrimaryKey("name", "rel_owner"),
	)
}

func (s *Store) MessagingAttachmentTableDef() *Table {
	// @todo merge with general attachment table
	return TableDef("messaging_attachment",
		AddID(),
		ColumnDef("rel_user", ColumnTypeIdentifier), // @todo rename => rel_owner
		ColumnDef("url", ColumnTypeText),
		ColumnDef("preview_url", ColumnTypeText),
		ColumnDef("name", ColumnTypeText),
		ColumnDef("meta", ColumnTypeJson),
		CUDTimestamps,
	)
}

func (s *Store) MessagingChannelTableDef() *Table {
	return TableDef("messaging_channel",
		AddID(),
		ColumnDef("name", ColumnTypeText),
		ColumnDef("topic", ColumnTypeText),
		ColumnDef("meta", ColumnTypeJson),
		ColumnDef("type", ColumnTypeText),
		ColumnDef("membership_policy", ColumnTypeText),
		ColumnDef("rel_creator", ColumnTypeIdentifier), // @todo rename => created_by
		ColumnDef("archived_at", ColumnTypeTimestamp),
		ColumnDef("rel_last_message", ColumnTypeIdentifier),
		CUDTimestamps,
	)
}

func (s *Store) MessagingChannelMemberTableDef() *Table {
	return TableDef("messaging_channel_member",
		ColumnDef("rel_channel", ColumnTypeIdentifier),
		ColumnDef("rel_user", ColumnTypeIdentifier),
		ColumnDef("type", ColumnTypeText),
		ColumnDef("flag", ColumnTypeText),
		CUDTimestamps,
		SetPrimaryKey("rel_channel", "rel_user"),
	)
}

func (s *Store) MessagingMentionTableDef() *Table {
	return TableDef("messaging_mention",
		ColumnDef("rel_channel", ColumnTypeIdentifier),
		ColumnDef("rel_message", ColumnTypeIdentifier),
		ColumnDef("rel_user", ColumnTypeIdentifier),
		ColumnDef("rel_mentioned_by", ColumnTypeIdentifier),
		ColumnDef("created_at", ColumnTypeTimestamp),
	)
}

func (s *Store) MessagingMessageTableDef() *Table {
	return TableDef("messaging_message",
		AddID(),
		ColumnDef("type", ColumnTypeText),
		ColumnDef("message", ColumnTypeText),
		ColumnDef("meta", ColumnTypeJson),
		ColumnDef("rel_channel", ColumnTypeIdentifier),
		ColumnDef("rel_user", ColumnTypeIdentifier),
		ColumnDef("reply_to", ColumnTypeIdentifier, DefaultValue("0")),
		ColumnDef("replies", ColumnTypeInteger, DefaultValue("0")),

		CUDTimestamps,
	)
}

func (s *Store) MessagingMessageAttachmentTableDef() *Table {
	return TableDef("messaging_message_attachment",
		ColumnDef("rel_message", ColumnTypeIdentifier),
		ColumnDef("rel_attachment", ColumnTypeIdentifier),
		SetPrimaryKey("rel_message"),
	)

}

func (s *Store) MessagingMessageFlagTableDef() *Table {
	return TableDef("messaging_message_flag",
		AddID(),
		ColumnDef("rel_channel", ColumnTypeIdentifier),
		ColumnDef("rel_message", ColumnTypeIdentifier),
		ColumnDef("rel_user", ColumnTypeIdentifier),
		ColumnDef("flag", ColumnTypeText),
		ColumnDef("created_at", ColumnTypeTimestamp),
	)
}

func (s *Store) MessagingPermissionRulesTableDef() *Table {
	return TableDef("messaging_permission_rules",
		ColumnDef("rel_role", ColumnTypeIdentifier),
		ColumnDef("resource", ColumnTypeText),
		ColumnDef("operation", ColumnTypeText),
		ColumnDef("access", ColumnTypeInteger),

		SetPrimaryKey("rel_role", "resource", "operation", "access"),
	)
}

func (s *Store) MessagingSettingsTableDef() *Table {
	return TableDef("messaging_settings", ColumnDef("rel_owner", ColumnTypeIdentifier),
		ColumnDef("name", ColumnTypeText, ColumnTypeLength(200)),
		ColumnDef("value", ColumnTypeJson),
		ColumnDef("updated_by", ColumnTypeIdentifier),
		ColumnDef("updated_at", ColumnTypeTimestamp),

		SetPrimaryKey("name", "rel_owner"),
	)

}

func (s *Store) MessagingUnreadTableDef() *Table {
	return TableDef("messaging_unread",
		ColumnDef("rel_channel", ColumnTypeIdentifier),
		ColumnDef("rel_reply_to", ColumnTypeIdentifier),
		ColumnDef("rel_user", ColumnTypeIdentifier),
		ColumnDef("count", ColumnTypeInteger),
		ColumnDef("rel_last_message", ColumnTypeIdentifier),
		SetPrimaryKey("rel_channel", "rel_reply_to", "rel_user"),
	)
}
