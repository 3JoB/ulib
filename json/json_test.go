package json_test

import (
	"testing"

	"github.com/3JoB/telebot/pkg"
	"github.com/3JoB/ulib/json"
)

type TestStruct struct{
	A string `json:"a"`
}

func TestMarshrl(t *testing.T) {
	data := `{"a": "b"}`
	var tsc TestStruct
	if err := json.Unmarshal(pkg.Bytes(data), &tsc); err!= nil {
		panic(err)
	}
	da := json.Marshal(&tsc).String()
	println(da)
}