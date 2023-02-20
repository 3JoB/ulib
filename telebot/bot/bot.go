package bot

import (
	"fmt"
	"net/http"
	"os"

	tele "github.com/3JoB/telebot"
	telemw "github.com/3JoB/telebot/middleware"

	"github.com/3JoB/ulib/telebot/middleware"
)

type bot struct {
	Settings tele.Settings
}

type tb struct {
	B *tele.Bot
}

// Initialize robot
func New() *bot {
	return new(bot)
}

/*
Example:

	&tele.Webhook{
			Endpoint:       &tele.WebhookEndpoint{PublicURL: webhookEndpoint},
			AllowedUpdates: []string{"callback_query", "message"},
			Listen:         ":8888",
		},
*/
func (b *bot) SetWebHook(webhook *tele.Webhook) *bot {
	b.Settings.Poller = webhook
	return b
}

// Set Robot APIKey
func (b *bot) SetKey(key string) *bot {
	b.Settings.Token = key
	return b
}

// Set error handling
func (b *bot) SetError(endpoint func(error, tele.Context)) *bot {
	b.Settings.OnError = endpoint
	return b
}

// Set Client (for net/http)
func (b *bot) SetClient(end *http.Client) *bot {
	b.Settings.Client = end
	return b
}

func (b *bot) SetUpdates(updates int) *bot {
	b.Settings.Updates = updates
	return b
}

// Custom Settings
func (b *bot) CustomSettings(settings tele.Settings) *bot {
	b.Settings = settings
	return b
}

// Initialize with configuration
func (b *bot) CreateBot() *tb {
	var err error
	t := new(tb)
	t.B, err = tele.NewBot(b.Settings)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-2)
	}
	return t
}

// Remove Webhooks
func (b *tb) RemoveWebhook() {
	b.B.RemoveWebhook(true)
}

// Import the middleware
func (b *tb) Middleware(m ...tele.MiddlewareFunc) {
	b.B.Use(m...)
}

// Import logging middleware (use default configuration)
func (b *tb) ImportMiddlewareLogger() {
	b.B.Use(middleware.Logger(nil))
}

// Import Recover middleware
func (b *tb) ImportMiddlewareRecover() {
	b.B.Use(telemw.Recover())
}

// Import handler
func (b *tb) Handle(endpoint any, h tele.HandlerFunc, m ...tele.MiddlewareFunc) {
	if len(m) != 0 {
		b.B.Handle(endpoint, h, m...)
	} else {
		b.B.Handle(endpoint, h)
	}
}

func (b *tb) Me() *tele.User {
	return b.B.Me
}

// Start the robot
func (b *tb) Start() {
	b.B.Start()
}
