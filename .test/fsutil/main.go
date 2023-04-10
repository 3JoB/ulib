package main

import (
	"time"

	"github.com/3JoB/ulib/fsutil"
)

func main() {
	f, _ := fsutil.Create("a.txt")
	f.Close()
	time.Sleep(time.Second * 5)
	fsutil.Remove("a.txt")
}
