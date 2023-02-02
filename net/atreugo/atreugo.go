package atreugo_utils

import (
	"github.com/3JoB/ulib/json"
	"github.com/savsgio/atreugo/v11"
)

func Bind(c *atreugo.RequestCtx, v any) {
	json.Unmarshal(c.Request.Body(), v)
}