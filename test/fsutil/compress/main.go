package main

import (
	"fmt"

	"github.com/3JoB/ulib/fsutil"
	"github.com/3JoB/ulib/fsutil/compress"
)

var Mode string

func main() {
	if Mode == "true" {
		compress.NewZip().Extract("test.zst.zip", "testios")
	} else {
		data := fsutil.ReadDirAll("testdata")
		if err := compress.NewZip().Create("test.zst.zip", data); err != nil {
			panic(err)
		}
		fmt.Println("ok!")
	}
}
