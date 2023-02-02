package telebot_test

import (
	"testing"

	tele "github.com/3JoB/telebot"

	"github.com/3JoB/ulib/telebot"
)

func TestSendMessage(t *testing.T) {
	var ctx tele.Context
	n := telebot.New().SetContext(ctx)
	n.SetChatID(114514).SetAutoDelete(15).Send("haha")
	// or
	n.SetAutoDelete(15).Send("haha")
	// or
	file := &tele.Video{
		File: tele.FromDisk("test.mp4"),
	}
	n.SetModes(telebot.ModeFile).Send(file)
}

func TestShowAlert(t *testing.T) {
	var ctx tele.Context
	n := telebot.New()
	n.SetContext(ctx).SetShowAlert().Alert("hello")
}

func TestSendMedia(t *testing.T) {}

func TestAll(t *testing.T) {
	var ctx tele.Context
	n := telebot.New()
	info := n.SetContext(ctx).SetChatID(114514)
	// send msg
	info.Send("nullcx")
	info.Alert("ccy")
}
