package client_utils

import (
	"io"
	"net/http"

	"github.com/3JoB/telebot/pkg"
	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/flate"
	"github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/zstd"
	"github.com/zc310/headers"
)

var decoder, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(0))

type update struct{
	data []byte
	Err error
}

func Data(r *http.Response) *update {
	upd := new(update)
	upd.data, upd.Err = upk(r)
	return upd
}

func (u *update) String() string {
	udt := pkg.String(u.data)
	return udt
}

func (u *update) Bytes() []byte {
	return u.data
}

func upk(r *http.Response) ([]byte, error) {
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