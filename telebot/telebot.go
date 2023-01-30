package telebot

import (
	"errors"
	"time"

	tele "github.com/3JoB/telebot"
	"github.com/3JoB/telebot/pkg"
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
	Err             error
}

var (
	SendOptions = &tele.SendOptions{ParseMode: tele.ModeHTML, DisableWebPagePreview: true}
)

func New() *use {
	n := new(use)
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
	return n.Ctx.Bot().Delete(&tele.StoredMessage{
		MessageID: cast.ToString(message),
		ChatID:    n.ChatId,
	})
}

// Send Message
func (n *use) Send(msg any) (*tele.Message, error) {
	var i *tele.Message
	if n.Btn != nil {
		i, n.Err = n.Ctx.Send(msg, SendOptions, n.Btn)
	} else {
		i, n.Err = n.Ctx.Send(msg, SendOptions)
	}
	if n.AutoDelete {
		time.Sleep(time.Second * n.AutoDeleteTimer)
		n.Delete(i.ID)
	}
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
	if n.ChatId == 0 {
		return nil, errors.New("chatid not available")
	}
	if n.Ctx == nil {
		return nil, errors.New("context not set")
	}
	var b []AdminList
	d, _ := n.Ctx.Bot().Raw("getChatAdministrators", map[string]int64{"chat_id": n.ChatId})
	if !gjson.GetBytes(d, "ok").Bool() {
		return nil, nil
	}
	json.Unmarshal(pkg.Bytes(gjson.GetBytes(d, "result").String()), &b)
	if len(b) == 0 {
		return nil, errors.New("failed to fetch admin list,please check what happened")
	}
	admin := make(map[int64]int)
	for _, i := range b {
		admin[i.User.ID] = 1
	}
	return admin, nil
}
