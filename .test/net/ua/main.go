package main

import (
	"fmt"

	"github.com/3JoB/ulib/net/ua"
)

func main() {
	u := ua.GenerateUA(ua.Config{
		SoftInfo: "ULIB/v1.32.0",
	})
	fmt.Println(u)
}
