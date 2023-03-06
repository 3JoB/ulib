package main

import (
	"fmt"

	"github.com/3JoB/ulib/fsutil"
)

func main() {
	fmt.Println(fsutil.IsFile("main.go"))
}
