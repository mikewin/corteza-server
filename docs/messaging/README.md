# User activity

| Method | Endpoint | Purpose |
| ------ | -------- | ------- |
| `POST` | `/activity/` | Sends user's activity to all subscribers; globally or per channel/message. |

## Sends user's activity to all subscribers; globally or per channel/message.

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/activity/` | HTTP/S | POST |  |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | uint64 | POST | Channel ID, if set, activity will be send only to subscribed users | N/A | NO |
| messageID | uint64 | POST | Message ID, if set, channelID must be set as well | N/A | NO |
| kind | string | POST | Arbitrary string | N/A | YES |

---




# Attachments

| Method | Endpoint | Purpose |
| ------ | -------- | ------- |
| `GET` | `/attachment/{attachmentID}/original/{name}` | Serves attached file |
| `GET` | `/attachment/{attachmentID}/preview.{ext}` | Serves preview of an attached file |

## Serves attached file

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/attachment/{attachmentID}/original/{name}` | HTTP/S | GET | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| download | bool | GET | Force file download | N/A | NO |
| sign | string | GET | Signature | N/A | YES |
| userID | uint64 | GET | User ID | N/A | YES |
| name | string | PATH | File name | N/A | YES |
| attachmentID | uint64 | PATH | Attachment ID | N/A | YES |

## Serves preview of an attached file

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/attachment/{attachmentID}/preview.{ext}` | HTTP/S | GET | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| ext | string | PATH | Preview extension/format | N/A | YES |
| attachmentID | uint64 | PATH | Attachment ID | N/A | YES |
| sign | string | GET | Signature | N/A | YES |
| userID | uint64 | GET | User ID | N/A | YES |

---




# Channels

A channel is a representation of a sequence of messages. It has meta data like channel subject. Channels may be public, private or group.

| Method | Endpoint | Purpose |
| ------ | -------- | ------- |
| `GET` | `/channels/` | List channels |
| `POST` | `/channels/` | Create new channel |
| `PUT` | `/channels/{channelID}` | Update channel details |
| `PUT` | `/channels/{channelID}/state` | Update channel state |
| `PUT` | `/channels/{channelID}/flag` | Update channel membership flag |
| `DELETE` | `/channels/{channelID}/flag` | Remove channel membership flag |
| `GET` | `/channels/{channelID}` | Read channel details |
| `GET` | `/channels/{channelID}/members` | List channel members |
| `PUT` | `/channels/{channelID}/members/{userID}` | Join channel |
| `DELETE` | `/channels/{channelID}/members/{userID}` | Remove member from channel |
| `POST` | `/channels/{channelID}/invite` | Join channel |
| `POST` | `/channels/{channelID}/attach` | Attach file to channel |

## List channels

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/` | HTTP/S | GET | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| query | string | GET | Search query | N/A | NO |

## Create new channel

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/` | HTTP/S | POST | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| name | string | POST | Name of Channel | N/A | NO |
| topic | string | POST | Subject of Channel | N/A | NO |
| type | string | POST | Channel type | N/A | NO |
| membershipPolicy | types.ChannelMembershipPolicy | POST | Membership policy (eg: featured, forced)? | N/A | NO |
| members | []string | POST | Initial members of the channel | N/A | NO |

## Update channel details

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}` | HTTP/S | PUT | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | uint64 | PATH | Channel ID | N/A | YES |
| name | string | POST | Name of Channel | N/A | NO |
| topic | string | POST | Subject of Channel | N/A | NO |
| membershipPolicy | types.ChannelMembershipPolicy | POST | Membership policy (eg: featured, forced)? | N/A | NO |
| type | string | POST | Channel type | N/A | NO |
| organisationID | uint64 | POST | Move channel to different organisation | N/A | NO |

## Update channel state

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/state` | HTTP/S | PUT | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | uint64 | PATH | Channel ID | N/A | YES |
| state | string | POST | Valid values: delete, undelete, archive, unarchive | N/A | YES |

## Update channel membership flag

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/flag` | HTTP/S | PUT | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | uint64 | PATH | Channel ID | N/A | YES |
| flag | string | POST | Valid values: pinned, hidden, ignored | N/A | YES |

## Remove channel membership flag

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/flag` | HTTP/S | DELETE | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

## Read channel details

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}` | HTTP/S | GET | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

## List channel members

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/members` | HTTP/S | GET | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

## Join channel

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/members/{userID}` | HTTP/S | PUT | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | uint64 | PATH | Channel ID | N/A | YES |
| userID | uint64 | PATH | Member ID | N/A | NO |

## Remove member from channel

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/members/{userID}` | HTTP/S | DELETE | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | uint64 | PATH | Channel ID | N/A | YES |
| userID | uint64 | PATH | Member ID | N/A | NO |

## Join channel

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/invite` | HTTP/S | POST | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | uint64 | PATH | Channel ID | N/A | YES |
| userID | []string | POST | User ID | N/A | NO |

## Attach file to channel

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/attach` | HTTP/S | POST | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | uint64 | PATH | Channel ID | N/A | YES |
| replyTo | uint64 | POST | Upload as a reply | N/A | NO |
| upload | *multipart.FileHeader | POST | File to upload | N/A | YES |

---




# Commands

| Method | Endpoint | Purpose |
| ------ | -------- | ------- |
| `GET` | `/commands/` | List of available commands |

## List of available commands

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/commands/` | HTTP/S | GET |  |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |

---




# Messages

| Method | Endpoint | Purpose |
| ------ | -------- | ------- |
| `POST` | `/channels/{channelID}/messages/` | Post new message to the channel |
| `POST` | `/channels/{channelID}/messages/command/{command}/exec` | Execute command |
| `GET` | `/channels/{channelID}/messages/mark-as-read` | Manages read/unread messages in a channel or a thread |
| `PUT` | `/channels/{channelID}/messages/{messageID}` | Edit existing message |
| `DELETE` | `/channels/{channelID}/messages/{messageID}` | Delete existing message |
| `POST` | `/channels/{channelID}/messages/{messageID}/replies` | Reply to a message |
| `POST` | `/channels/{channelID}/messages/{messageID}/pin` | Pin message to channel (public bookmark) |
| `DELETE` | `/channels/{channelID}/messages/{messageID}/pin` | Pin message to channel (public bookmark) |
| `POST` | `/channels/{channelID}/messages/{messageID}/bookmark` | Bookmark a message (private bookmark) |
| `DELETE` | `/channels/{channelID}/messages/{messageID}/bookmark` | Remove boomark from message (private bookmark) |
| `POST` | `/channels/{channelID}/messages/{messageID}/reaction/{reaction}` | React to a message |
| `DELETE` | `/channels/{channelID}/messages/{messageID}/reaction/{reaction}` | Delete reaction from a message |

## Post new message to the channel

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/` | HTTP/S | POST | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| message | string | POST | Message contents (markdown) | N/A | YES |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

## Execute command

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/command/{command}/exec` | HTTP/S | POST | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| command | string | PATH | Command to be executed | N/A | YES |
| channelID | uint64 | PATH | Channel ID | N/A | YES |
| input | string | POST | Arbitrary command input | N/A | NO |
| params | []string | POST | Command parameters | N/A | NO |

## Manages read/unread messages in a channel or a thread

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/mark-as-read` | HTTP/S | GET | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| threadID | uint64 | GET | ID of thread (messageID)  | N/A | NO |
| lastReadMessageID | uint64 | GET | ID of the last read message | N/A | NO |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

## Edit existing message

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/{messageID}` | HTTP/S | PUT | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| messageID | uint64 | PATH | Message ID | N/A | YES |
| channelID | uint64 | PATH | Channel ID | N/A | YES |
| message | string | POST | Message contents (markdown) | N/A | YES |

## Delete existing message

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/{messageID}` | HTTP/S | DELETE | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| messageID | uint64 | PATH | Message ID | N/A | YES |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

## Reply to a message

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/{messageID}/replies` | HTTP/S | POST | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| messageID | uint64 | PATH | Message ID | N/A | YES |
| channelID | uint64 | PATH | Channel ID | N/A | YES |
| message | string | POST | Message contents (markdown) | N/A | YES |

## Pin message to channel (public bookmark)

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/{messageID}/pin` | HTTP/S | POST | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| messageID | uint64 | PATH | Message ID | N/A | YES |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

## Pin message to channel (public bookmark)

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/{messageID}/pin` | HTTP/S | DELETE | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| messageID | uint64 | PATH | Message ID | N/A | YES |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

## Bookmark a message (private bookmark)

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/{messageID}/bookmark` | HTTP/S | POST | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| messageID | uint64 | PATH | Message ID | N/A | YES |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

## Remove boomark from message (private bookmark)

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/{messageID}/bookmark` | HTTP/S | DELETE | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| messageID | uint64 | PATH | Message ID | N/A | YES |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

## React to a message

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/{messageID}/reaction/{reaction}` | HTTP/S | POST | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| messageID | uint64 | PATH | Message ID | N/A | YES |
| reaction | string | PATH | Reaction | N/A | YES |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

## Delete reaction from a message

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/channels/{channelID}/messages/{messageID}/reaction/{reaction}` | HTTP/S | DELETE | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| messageID | uint64 | PATH | Message ID | N/A | YES |
| reaction | string | PATH | Reaction | N/A | YES |
| channelID | uint64 | PATH | Channel ID | N/A | YES |

---




# Permissions

| Method | Endpoint | Purpose |
| ------ | -------- | ------- |
| `GET` | `/permissions/` | Retrieve defined permissions |
| `GET` | `/permissions/effective` | Effective rules for current user |
| `GET` | `/permissions/{roleID}/rules` | Retrieve role permissions |
| `DELETE` | `/permissions/{roleID}/rules` | Remove all defined role permissions |
| `PATCH` | `/permissions/{roleID}/rules` | Update permission settings |

## Retrieve defined permissions

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/permissions/` | HTTP/S | GET | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |

## Effective rules for current user

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/permissions/effective` | HTTP/S | GET | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| resource | string | GET | Show only rules for a specific resource | N/A | NO |

## Retrieve role permissions

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/permissions/{roleID}/rules` | HTTP/S | GET | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| roleID | uint64 | PATH | Role ID | N/A | YES |

## Remove all defined role permissions

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/permissions/{roleID}/rules` | HTTP/S | DELETE | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| roleID | uint64 | PATH | Role ID | N/A | YES |

## Update permission settings

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/permissions/{roleID}/rules` | HTTP/S | PATCH | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| roleID | uint64 | PATH | Role ID | N/A | YES |
| rules | permissions.RuleSet | POST | List of permission rules to set | N/A | YES |

---




# Search entry point

| Method | Endpoint | Purpose |
| ------ | -------- | ------- |
| `GET` | `/search/messages` | Search for messages |
| `GET` | `/search/threads` | Search for threads |

## Search for messages

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/search/messages` | HTTP/S | GET | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | []string | GET | Filter by channels | N/A | NO |
| afterMessageID | uint64 | GET | ID of the first message in the list (exclusive) | N/A | NO |
| beforeMessageID | uint64 | GET | ID of the last message in the list (exclusive) | N/A | NO |
| fromMessageID | uint64 | GET | ID of the first message in the list (inclusive) | N/A | NO |
| toMessageID | uint64 | GET | ID of the last message the list (inclusive) | N/A | NO |
| threadID | []string | GET | Filter by thread message ID | N/A | NO |
| userID | []string | GET | Filter by one or more user | N/A | NO |
| type | []string | GET | Filter by message type (text, inlineImage, attachment, ...) | N/A | NO |
| pinnedOnly | bool | GET | Return only pinned messages | N/A | NO |
| bookmarkedOnly | bool | GET | Only bookmarked messages | N/A | NO |
| limit | uint | GET | Max number of messages | N/A | NO |
| query | string | GET | Search query | N/A | NO |

## Search for threads

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/search/threads` | HTTP/S | GET | Client ID, Session ID |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| channelID | []string | GET | Filter by channels | N/A | NO |
| limit | uint | GET | Max number of messages | N/A | NO |
| query | string | GET | Search query | N/A | NO |

---




# Settings

| Method | Endpoint | Purpose |
| ------ | -------- | ------- |
| `GET` | `/settings/` | List settings |
| `PATCH` | `/settings/` | Update settings |
| `GET` | `/settings/{key}` | Get a value for a key |
| `GET` | `/settings/current` | Current compose settings |

## List settings

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/settings/` | HTTP/S | GET |  |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| prefix | string | GET | Key prefix | N/A | NO |

## Update settings

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/settings/` | HTTP/S | PATCH |  |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| values | settings.ValueSet | POST | Array of new settings: `[{ name: ..., value: ... }]`. Omit value to remove setting | N/A | YES |

## Get a value for a key

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/settings/{key}` | HTTP/S | GET |  |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| ownerID | uint64 | GET | Owner ID | N/A | NO |
| key | string | PATH | Setting key | N/A | YES |

## Current compose settings

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/settings/current` | HTTP/S | GET |  |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |

---




# Status

| Method | Endpoint | Purpose |
| ------ | -------- | ------- |
| `GET` | `/status/` | See all current statuses |
| `POST` | `/status/` | Set user's status |
| `DELETE` | `/status/` | Clear status |

## See all current statuses

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/status/` | HTTP/S | GET |  |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |

## Set user's status

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/status/` | HTTP/S | POST |  |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |
| icon | string | POST | Status icon | N/A | NO |
| message | string | POST | Status message | N/A | NO |
| expires | string | POST | Clear status when it expires (eg: when-active, afternoon, tomorrow 1h, 30m, 1 PM, 2019-05-20) | N/A | NO |

## Clear status

#### Method

| URI | Protocol | Method | Authentication |
| --- | -------- | ------ | -------------- |
| `/status/` | HTTP/S | DELETE |  |

#### Request parameters

| Parameter | Type | Method | Description | Default | Required? |
| --------- | ---- | ------ | ----------- | ------- | --------- |

---