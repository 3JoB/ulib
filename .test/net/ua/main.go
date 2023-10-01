package main

import (
	"fmt"

	"github.com/3JoB/ulib/net/ua"
)

func main() {
	u := ua.GenerateUA(ua.Config{
		SoftInfo: "ULIB/v1.36.2",
	})
	fmt.Println(u)
}
