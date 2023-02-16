package json_test

import (
	"fmt"
	"testing"

	"github.com/3JoB/ulib/json"
)

type TestStruct struct {
	A string `json:"a"`
}

func TestMain(t *testing.T) {
	data := `{"a": "b"}`
	var tsc TestStruct
	if err := json.UnmarshalString(data, &tsc); err != nil {
		panic(err)
	}
	da := json.Marshal(&tsc).String()
	fmt.Println(da)
}
