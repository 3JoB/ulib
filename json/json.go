package json

import (
	"github.com/3JoB/telebot/pkg"
	gjs "github.com/goccy/go-json"
)

type marshal struct {
	err  error
	data []byte
}

func Marshal(a any) *marshal {
	m := new(marshal)
	m.data, m.err = gjs.Marshal(a)
	return m
}

func Unmarshal(data []byte, str any) error {
	return gjs.Unmarshal(data, str)
}

func (m *marshal) String(a any) string {
	return pkg.String(m.data)
}
