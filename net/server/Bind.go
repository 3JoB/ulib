package server

import (
	"io"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/json"
)

/*
gin:
server.Bind(c.Request.Body, &v)
*/
func Bind(r io.ReadCloser, v any) {
	data, _ := io.ReadAll(r)
	json.Unmarshal(data, v)
}

func Body(r io.ReadCloser, l int64) string {
	body := make([]byte, l)
	r.Read(body)
	return unsafeConvert.String(body)
}
