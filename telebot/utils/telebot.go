package utils

import (
	"errors"
	"time"

	tele "github.com/3JoB/telebot"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"

	"github.com/3JoB/ulib/json"
)

type Use struct {
	Context         tele.Context
	ChatId          int64
	AutoDeleteTimer time.Duration
	AutoDelete      bool
	ShowAlert       bool
	DeleteCommand   bool
	Threads         bool
	ParseMode       string
	Btn             *tele.ReplyMarkup
	SendOptions     *tele.SendOptions
}

type SendMode string

const (
	ModeDefault SendMode = SendMode(tele.ModeDefault)
	ModeHTML    SendMode = SendMode(tele.ModeHTML)
	ModeMD      SendMode = SendMode(tele.ModeMarkdown)
	ModeMD2     SendMode = SendMode(tele.ModeMarkdownV2)
	ModeFile    SendMode = SendMode("File")
)

var (
	ErrCtxNotSet    = errors.New("ulib.telebot: context not set")
	ErrNoSuperGroup = errors.New("ulib.telebot: Cannot be non-supergroup type")
)

func New() *Use {
	n := new(Use)
	n.SendOptions, n.Btn = new(tele.SendOptions), new(tele.ReplyMarkup)
	return n
}

func (n *Use) ImportOpts(v *tele.SendOptions) *Use {
	n.SendOptions = v
	return nil
}

// Leave makes bot leave a group, supergroup or channel.
func (n *Use) Leave(v ...*tele.Chat) error {
	if len(v) == 0 {
		return n.Context.Bot().Leave(n.Context.Chat())
	}
	return n.Context.Bot().Leave(v[0])
}

// Delete Message
func (n *Use) Delete(message int) error {
	if n.Context == nil {
		return ErrCtxNotSet
	}

	var c int64

	if n.ChatId != 0 {
		c = n.ChatId
	} else {
		c = n.Context.Chat().ID
	}

	return n.Context.Bot().Delete(&tele.StoredMessage{
		MessageID: cast.ToString(message),
		ChatID:    c,
	})
}

// Send Message
func (n *Use) Send(v any) (i *tele.Message, e error) {
	var c tele.ChatID

	if n.Context == nil {
		return nil, ErrCtxNotSet
	}

	if n.ChatId != 0 {
		c = tele.ChatID(n.ChatId)
	} else {
		c = tele.ChatID(n.Context.Chat().ID)
	}

	if n.Context.Chat().IsForum {
		if !n.Threads {
			n.SendOptions.Thread = &tele.Topic{
				ThreadID: cast.ToInt64(n.Context.Message().ThreadID),
			}
		}
	} else {
		if n.SendOptions.Thread != nil {
			n.SendOptions.Thread = nil
		}
	}

	if n.Btn != nil {
		i, e = n.Context.Bot().Send(c, v, n.SendOptions, n.Btn)
	} else {
		i, e = n.Context.Bot().Send(c, v, n.SendOptions)
	}

	if e != nil {
		return i, e
	}

	if n.DeleteCommand {
		n.Context.Delete()
	}

	if n.AutoDelete {
		time.Sleep(time.Second * n.AutoDeleteTimer)
		if err := n.Delete(i.ID); err != nil {
			return nil, err
		}
	}

	n.AutoDelete = false

	return i, e
}

// Pop-ups
func (n *Use) Alert(text string) error {
	if n.Context == nil {
		return ErrCtxNotSet
	}

	return n.Context.Respond(&tele.CallbackResponse{
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
