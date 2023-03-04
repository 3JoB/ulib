package main

import (
	"fmt"
	"io"
	"os"

	"github.com/3JoB/ulib/fsutil"
	"github.com/3JoB/ulib/net/client"
	"github.com/go-resty/resty/v2"
)

func main() {
	r, _ := resty.New().R().Get("https:example.com/example.json")
	data := client.UnPackData(r.RawResponse).String()
	fmt.Println(data)

	ds := client.GetSource(r.RawResponse)
	file, _ := fsutil.OpenFile("main.json", os.O_RDWR, 0755)
	switch dse := ds.(type) {
	case io.ReadCloser:
		io.Copy(file, dse)
		dse.Close()
		return
	case io.Reader:
		io.Copy(file, dse)
		return
	case error:
		fmt.Println(dse.Error())
		return
	default:
		return
	}
}