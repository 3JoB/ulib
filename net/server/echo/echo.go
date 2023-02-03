package echo_utils

import (
	"io"

	"github.com/3JoB/telebot/pkg"
	"github.com/labstack/echo/v4"

	"github.com/3JoB/ulib/json"
)

func Bind(c echo.Context, v any) {
	data, _ := io.ReadAll(c.Request().Body)
	json.Unmarshal(data, v)
}

func Body(c echo.Context) string {
	body := make([]byte, c.Request().ContentLength)
	c.Request().Body.Read(body)
	return pkg.String(body)
}
