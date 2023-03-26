package main

import (
	"fmt"

	"github.com/3JoB/ulib/fsutil"
	"github.com/3JoB/ulib/fsutil/compress/zszip"
)

var Mode string

func main() {
	if Mode == "true" {
		zszip.New().Extract("test.zst.zip", "testios")
	} else {
		data := fsutil.ReadDirAll("testdata")
		if err := zszip.New().Create("test.zst.zip", data); err != nil {
			panic(err)
		}
		fmt.Println("ok!")
	}
}
