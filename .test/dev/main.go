package main

import (
	iocopy "github.com/3JoB/ulib/dev/ioCopy"
	"github.com/3JoB/ulib/fsutil"
)

func main() {
	t1, _ := fsutil.Open("test.dev")
	t2, _ := fsutil.Open("test2.dev")
	_, err := iocopy.Copy(t1, t2)
	if err != nil {
		panic(err)
	}
}
