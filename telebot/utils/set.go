package utils

import (
	"time"

	tele "github.com/3JoB/telebot"
)

func (n *Use) SetModes(s ...SendMode) *Use {
	if len(s) == 0 {
		return n
	}
	switch s[0] {
	case ModeDefault:
		n.SendOptions.ParseMode = tele.ModeHTML
	case ModeHTML:
		n.SendOptions.ParseMode = tele.ModeHTML
	case ModeMD:
		n.SendOptions.ParseMode = tele.ModeMarkdown
	case ModeMD2:
		n.SendOptions.ParseMode = tele.ModeMarkdownV2
	case ModeFile:
		n.SendOptions.ParseMode = tele.ModeHTML
	default:
		n.SendOptions.ParseMode = tele.ModeHTML
	}
	n.ParseMode = n.SendOptions.ParseMode
	return n
}

// For text messages, disables previews for links in this message.
func (n *Use) SetDisableWebPreview() *Use {
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

func (n *Use) SetTopicID(v int64) *Use {
	n.Threads, n.SendOptions.Thread = true, &tele.Topic{ThreadID: v}
	return n
}

func (n *Use) SetContext(c tele.Context) *Use {
	n.Context = c
	return n
}

func (n *Use) SetAutoDelete(t time.Duration) *Use {
	n.AutoDelete, n.AutoDeleteTimer = true, t
	return n
}

func (n *Use) SetDeleteCommand() *Use {
	if n.DeleteCommand {
		n.DeleteCommand = false
	} else {
		n.DeleteCommand = true
	}
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
