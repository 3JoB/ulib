package utils

import (
	"errors"

	"github.com/spf13/cast"
	"github.com/tidwall/gjson"

	"github.com/3JoB/ulib/json"
)

type AdminInfo struct {
	User   User   `json:"user"`
	Status string `json:"status"`
	AdminPerm
}

type AdminPerm struct {
	CanBeEdited         bool `json:"can_be_edited"`
	CanManageChat       bool `json:"can_manage_chat"`
	CanChangeInfo       bool `json:"can_change_info"`
	CanDeleteMessages   bool `json:"can_delete_messages"`
	CanInviteUsers      bool `json:"can_invite_users"`
	CanRestrictMembers  bool `json:"can_restrict_members"`
	CanPinMessages      bool `json:"can_pin_messages"`
	CanManageTopics     bool `json:"can_manage_topics"`
	CanPromoteMembers   bool `json:"can_promote_members"`
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	IsAnonymous         bool `json:"is_anonymous"`
	CanManageVoiceChats bool `json:"can_manage_voice_chats"`
}

type User struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	IsPremium bool   `json:"is_premium"`
	Language  string `json:"language_code"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// GetChatMember Status
//
// 0 error, 1 Not in Chat, 2 In Chat, 3 Banned
func (n *Use) GetChatMember(uid int64) (int, error) {
	if uid == 0 {
		return 0, errors.New("uid is 0")
	}
	if n.ChatId == 0 {
		if n.Context.Chat().Type == "private" {
			return 0, errors.New("chat is private")
		}
		n.ChatId = n.Context.Chat().ID
	}
	d, err := n.Context.Bot().Raw("getChatMember", map[string]int64{"chat_id": n.ChatId, "user_id": uid})
	if err != nil {
		return 0, err
	}
	switch gjson.GetBytes(d, "result.status").String() {
	case "kicked":
		return 3, nil
	case "left":
		return 1, nil
	default:
		return 2, nil
	}
}

// Get the list of group administrators
//
// It is now recommended to use `telebot.Context.Bot().Adminsof()` method.
func (n *Use) GetAdminList() (map[int64]AdminInfo, error) {
	if n.Context == nil {
		return nil, ErrCtxNotSet
	}

	var b []AdminInfo
	var args map[string]int64 = map[string]int64{}

	if n.ChatId != 0 {
		args = map[string]int64{"chat_id": n.ChatId}
	} else {
		if n.Context.Chat().Type != "supergroup" {
			return nil, ErrNoSuperGroup
		}
		args = map[string]int64{"chat_id": n.Context.Chat().ID}
	}

	d, _ := n.Context.Bot().Raw("getChatAdministrators", args)

	if !gjson.GetBytes(d, "ok").Bool() {
		return nil, errors.New("ulib.telebot: failed to fetch admin list\ndata: " + cast.ToString(d))
	}

	json.UnmarshalString(gjson.GetBytes(d, "result").String(), &b)
	if len(b) == 0 {
		return nil, errors.New("ulib.telebot: failed to fetch admin list\ndata: " + cast.ToString(d))
	}

	admin := make(map[int64]AdminInfo)
	for _, i := range b {
		admin[i.User.ID] = i
	}
	return admin, nil
}
