package http_utils

import (
	"io"
	"net/http"

	"github.com/3JoB/telebot/pkg"
	"github.com/3JoB/ulib/json"
)

func Bind(c *http.Request, v any){
	data, _ := io.ReadAll(c.Body)
	json.Unmarshal(data, v)
}

func Body(c *http.Request) string {
	body := make([]byte, c.ContentLength)
	c.Body.Read(body)
	return pkg.String(body)
}