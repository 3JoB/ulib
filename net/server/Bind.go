package server

import (
	"io"

	"github.com/3JoB/unsafeConvert"
	"github.com/sugawarayuuta/sonnet"
)

/*
gin:
server.Bind(c.Request.Body, &v)
*/
func Bind(r io.ReadCloser, v any) {
	data, _ := io.ReadAll(r)
	sonnet.Unmarshal(data, v)
}

func Body(r io.ReadCloser, l int64) string {
	body := make([]byte, l)
	r.Read(body)
	return unsafeConvert.StringPointer(body)
}
