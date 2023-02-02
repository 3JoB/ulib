package atreugo_utils

import (
	"github.com/savsgio/atreugo/v11"

	"github.com/3JoB/ulib/json"
)

func Bind(c *atreugo.RequestCtx, v any) {
	json.Unmarshal(c.Request.Body(), v)
}
