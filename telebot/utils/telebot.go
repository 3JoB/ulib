package utils

import (
	"errors"
	"fmt"
	"time"

	tele "github.com/3JoB/telebot"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"

	"github.com/3JoB/ulib/json"
	"github.com/3JoB/ulib/reflect"
)

type Use struct {
	Ctx             tele.Context
	ChatId          int64
	AutoDeleteTimer time.Duration
	AutoDelete      bool
	ShowAlert       bool
	Btn             *tele.ReplyMarkup
	SendOptions     *tele.SendOptions
	FileMode        bool
	Err             error
}

const (
	ModeDef         = tele.ModeDefault
	ModeFile string = "File"
	ModeHTML        = tele.ModeHTML
	ModeMD          = tele.ModeMarkdown
	ModeMD2         = tele.ModeMarkdownV2
)

func New() *Use {
	n := new(Use)
	n.SendOptions = new(tele.SendOptions)
	n.Btn = new(tele.ReplyMarkup)
	return n
}

func (n *Use) SetModes(s any) *Use {
	switch cast.ToString(s) {
	case ModeDef:
		n.SendOptions.ParseMode = ModeDef
	case ModeHTML:
		n.SendOptions.ParseMode = ModeHTML
	case ModeMD:
		n.SendOptions.ParseMode = ModeMD
	case ModeMD2:
		n.SendOptions.ParseMode = ModeMD2
	case ModeFile:
		n.FileMode = true
	default:
		panic("ulib.telebot: Unavailable sending method!!!")
	}
	return n
}

func (n *Use) SetWebPreview() *Use {
	if n.SendOptions.DisableWebPagePreview {
		n.SendOptions.DisableWebPagePreview = false
	} else {
		n.SendOptions.DisableWebPagePreview = true
	}
	return n
}

func (n *Use) SetChatID(c int64) *Use {
	n.ChatId = c
	return n
}

func (n *Use) SetContext(c tele.Context) *Use {
	n.Ctx = c
	return n
}

func (n *Use) SetAutoDelete(t time.Duration) *Use {
	n.AutoDelete = true
	n.AutoDeleteTimer = t
	return n
}

func (n *Use) SetShowAlert() *Use {
	if n.ShowAlert {
		n.ShowAlert = false
	} else {
		n.ShowAlert = true
	}
	return n
}

func (n *Use) SetBtn(btn *tele.ReplyMarkup) *Use {
	n.Btn = btn
	return n
}

// Delete Message
func (n *Use) Delete(message int) error {
	var chid int64
	if n.ChatId != 0 {
		chid = n.ChatId
	} else {
		chid = n.Ctx.Chat().ID
	}
	return n.Ctx.Bot().Delete(&tele.StoredMessage{
		MessageID: cast.ToString(message),
		ChatID:    chid,
	})
}

// Send Message
func (n *Use) Send(msg any) (*tele.Message, error) {
	var i *tele.Message
	var cid tele.ChatID
	if n.SendOptions.ParseMode == "" {
		n.SendOptions.ParseMode = ModeHTML
	}
	if n.ChatId != 0 {
		cid = tele.ChatID(n.ChatId)
	} else {
		cid = tele.ChatID(n.Ctx.Chat().ID)
	}
	if n.FileMode {
		if n.Btn != nil {
			i, n.Err = n.Ctx.Bot().Send(cid, msg, n.Btn)
		} else {
			i, n.Err = n.Ctx.Bot().Send(cid, msg)
		}
	} else {
		if n.Btn != nil {
			i, n.Err = n.Ctx.Bot().Send(cid, msg, n.SendOptions, n.Btn)
		} else {
			i, n.Err = n.Ctx.Bot().Send(cid, msg, n.SendOptions)
		}
	}
	if n.AutoDelete {
		time.Sleep(time.Second * n.AutoDeleteTimer)
		n.Delete(i.ID)
	}
	n.AutoDelete = false
	n.FileMode = false

	return i, n.Err
}

// Pop-ups
func (n *Use) Alert(text string) error {
	return n.Ctx.Respond(&tele.CallbackResponse{
		Text:      text,
		ShowAlert: n.ShowAlert,
	})
}

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
	if n.Ctx == nil {
		return nil, errors.New("ulib.telebot: context not set")
	}

	var (
		b []AdminInfo
		d []byte
	)

	if n.ChatId != 0 {
		d, _ = n.Ctx.Bot().Raw("getChatAdministrators", map[string]int64{"chat_id": n.ChatId})
	} else {
		if n.Ctx.Chat().Type != "supergroup" {
			return nil, errors.New("ulib.telebot: Cannot be non-supergroup type")
		}
		d, _ = n.Ctx.Bot().Raw("getChatAdministrators", map[string]int64{"chat_id": n.Ctx.Chat().ID})
	}

	if !gjson.GetBytes(d, "ok").Bool() {
		return nil, nil
	}

	fmt.Println(reflect.String(d))

	json.UnmarshalString(gjson.GetBytes(d, "result").String(), &b)
	if len(b) == 0 {
		return nil, errors.New("ulib.telebot: failed to fetch admin list,please check what happened")
	}
	admin := make(map[int64]AdminInfo)
	for _, i := range b {
		admin[i.User.ID] = i
	}
	return admin, nil
}
