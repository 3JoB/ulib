package utils

import (
	"errors"
	"fmt"
	"github.com/3JoB/ulib/reflect"
	"time"

	tele "github.com/3JoB/telebot"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"

	"github.com/3JoB/ulib/json"
)

type AdminList struct {
	User struct {
		ID int64 `json:"id"`
	} `json:"user"`
}

type use struct {
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

func New() *use {
	n := new(use)
	n.SendOptions = new(tele.SendOptions)
	n.Btn = new(tele.ReplyMarkup)
	return n
}

func (n *use) SetModes(s any) *use {
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

func (n *use) SetWebPreview() *use {
	if n.SendOptions.DisableWebPagePreview {
		n.SendOptions.DisableWebPagePreview = false
	} else {
		n.SendOptions.DisableWebPagePreview = true
	}
	return n
}

func (n *use) SetChatID(c int64) *use {
	n.ChatId = c
	return n
}

func (n *use) SetContext(c tele.Context) *use {
	n.Ctx = c
	return n
}

func (n *use) SetAutoDelete(t time.Duration) *use {
	n.AutoDelete = true
	n.AutoDeleteTimer = t
	return n
}

func (n *use) SetShowAlert() *use {
	if n.ShowAlert {
		n.ShowAlert = false
	} else {
		n.ShowAlert = true
	}
	return n
}

func (n *use) SetBtn(btn *tele.ReplyMarkup) *use {
	n.Btn = btn
	return n
}

// Delete Message
func (n *use) Delete(message int) error {
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
func (n *use) Send(msg any) (*tele.Message, error) {
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
func (n *use) Alert(text string) error {
	return n.Ctx.Respond(&tele.CallbackResponse{
		Text:      text,
		ShowAlert: n.ShowAlert,
	})
}

// Get the list of group administrators
func (n *use) GetAdminList() (map[int64]int, error) {
	if n.Ctx == nil {
		return nil, errors.New("ulib.telebot: context not set")
	}

	var (
		b []AdminList
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
	admin := make(map[int64]int)
	for _, i := range b {
		admin[i.User.ID] = 1
	}
	return admin, nil
}
