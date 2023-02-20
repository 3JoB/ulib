package main

import (
	"fmt"

	"github.com/3JoB/ulib/fsutil"
	"github.com/3JoB/ulib/fsutil/compress"
)

func main(){
	data := fsutil.ReadPath("testdata")
	if err := compress.NewZip().Create("test.zst.zip", data); err != nil {
		panic(err)
	}
	fmt.Println("ok!")
}