// ulib client unpack data
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/3JoB/ulib/net/client"
//		"github.com/go-resty/resty/v2"
//	)
//
//	func main() {
//		r, err := resty.New().R().Get("https://example.com/example.json")
//		data := client.UnPackData(r.RawResponse).String()
//		fmt.Println(data)
//	}
package client

import (
	"io"
	"net/http"

	"github.com/3JoB/unsafeConvert"
	"github.com/JNyaa/headers"
	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/flate"
	"github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/zstd"

	"github.com/3JoB/ulib/json"
)

var decoder, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(0))

type update struct {
	data []byte
	err  error
}

// Decompress the Body package, support gzip, br, zstd, deflate
func UnPackData(r *http.Response) *update {
	u := &update{}
	u.unpack(r)
	return u
}

func GetSource(r *http.Response) any {
	switch r.Header.Get(headers.ContentEncoding) {
	case "br":
		return brotli.NewReader(r.Body)
	case "gzip":
		reader, err := gzip.NewReader(r.Body)
		if err != nil {
			return err
		}
		return reader
	case "zstd":
		return nil
	case "deflate":
		return flate.NewReader(r.Body)
	default:
		return nil
	}
}

// Return string type data
func (u *update) String() string {
	return unsafeConvert.String(u.data)
}

// Return []byte type data
func (u *update) Bytes() []byte {
	return u.data
}

// Return error data
func (u *update) Error() error {
	return u.err
}

// Directly bind the structure
func (u *update) Bind(v any) error {
	return json.Unmarshal(u.data, v)
}

func (u *update) unpack(r *http.Response) ([]byte, error) {
	switch r.Header.Get(headers.ContentEncoding) {
	case "br":
		reader := brotli.NewReader(r.Body)
		return io.ReadAll(reader)
	case "gzip":
		reader, err := gzip.NewReader(r.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		return io.ReadAll(reader)
	case "zstd":
		reader, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		return decoder.DecodeAll(reader, nil)
	case "deflate":
		reader := flate.NewReader(r.Body)
		defer reader.Close()
		return io.ReadAll(reader)
	default:
		return io.ReadAll(r.Body)
	}
}
