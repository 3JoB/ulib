package main

import (
	"github.com/3JoB/ulib/net/ua"
	"fmt"
)

func main() {
	u := ua.GenerateUA(ua.Config{
		SoftInfo: "ULIB/v1.4.0-Alpha",
	})
	fmt.Println(u)
}