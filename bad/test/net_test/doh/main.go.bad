package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/3JoB/ulib/bad/net/client/doh"
)

func main() {
	dialer, err := doh.Dialer(doh.Alicloud)
	if err != nil {
		panic(err)
	}
	httpClient := http.DefaultClient
	httpClient.Transport = http.DefaultTransport
	httpClient.Transport.(*http.Transport).DialContext = dialer.DialContext

	resp, err := httpClient.Get("https://lcag.org/gmake2.raw")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
