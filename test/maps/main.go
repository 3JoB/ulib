package main

import (
	"fmt"

	"github.com/3JoB/ulib/maps"
)

func main() {
	var m map[string]string
	m = maps.New(m)
	m["a"] = "b"
	fmt.Println(m)
}
