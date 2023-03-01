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

// Get the list of group administrators
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
