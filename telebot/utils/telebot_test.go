package utils_test

import (
	"testing"

	tele "github.com/3JoB/telebot"

	"github.com/3JoB/ulib/telebot/utils"
)

func TestSendMessage(t *testing.T) {
	var ctx tele.Context
	n := utils.New().SetContext(ctx)
	n.SetChatID(114514).SetAutoDelete(15).Send("haha")
	// or
	n.SetAutoDelete(15).Send("haha")
	// or
	file := &tele.Video{
		File: tele.FromDisk("test.mp4"),
	}
	n.Send(file)
}

func TestShowAlert(t *testing.T) {
	var ctx tele.Context
	n := utils.New()
	n.SetContext(ctx).SetShowAlert().Alert("hello")
}

func TestSendMedia(t *testing.T) {}

func TestAll(t *testing.T) {
	var ctx tele.Context
	n := utils.New()
	info := n.SetContext(ctx).SetChatID(114514)
	// send msg
	info.Send("nullcx")
	info.Alert("ccy")
}
