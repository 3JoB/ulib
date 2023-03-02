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
//		r, err := resty.New().Get("https://example.com/example.json")
//		data := client.UnPackData(r.RawBody()).String()
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
	u.data, u.err = unpack(r)
	return u
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

func unpack(r *http.Response) ([]byte, error) {
	switch r.Header.Get(headers.ContentEncoding) {
	case "br":
		return io.ReadAll(brotli.NewReader(r.Body))
	case "gzip":
		gr, err := gzip.NewReader(r.Body)
		if err != nil {
			return nil, err
		}
		return io.ReadAll(gr)
	case "zstd":
		reader, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		return decoder.DecodeAll(reader, nil)
	case "deflate":
		zr := flate.NewReader(r.Body)
		defer zr.Close()
		return io.ReadAll(zr)
	default:
		return io.ReadAll(r.Body)
	}
}
