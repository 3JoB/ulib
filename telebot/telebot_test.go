package telebot_test

import (
	tele "github.com/3JoB/telebot"
	"github.com/3JoB/ulib/telebot"
)

func SendMessage(){
	var ctx tele.Context
	n := telebot.New()
	n.SetContext(ctx).SetChatID(114514).SetAutoDelete(15).Send("haha")
}

func ShowAlert(){
	var ctx tele.Context
	n := telebot.New()
	n.SetContext(ctx).SetChatID(114514).SetShowAlert().Alert("hello")
}

func All(){
	var ctx tele.Context
	n := telebot.New()
	info := n.SetContext(ctx).SetChatID(114514)
	//send msg
	info.Send("nullcx")
	info.Alert("ccy")
}