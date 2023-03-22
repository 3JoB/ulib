package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/3JoB/ulib/net/client"
)

func main() {
	r, _ := resty.New().R().Get("https://example.com/example.json")
	data := client.UnPackData(r.RawResponse).String()
	fmt.Println(data)
}
