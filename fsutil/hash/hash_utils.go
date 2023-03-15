package hash

import (
	"hash"
	"io"

	"github.com/3JoB/ulib/fsutil"
	"github.com/3JoB/ulib/hex"
)

func hashel(path string, hs hash.Hash) string {
	f, err := fsutil.Open(path)
	if err != nil {
		f.Close()
		return ""
	}
	defer f.Close()
	_, _ = io.Copy(hs, f)
	return hex.EncodeToString(hs.Sum(nil))
}

func hash32el(path string, hs hash.Hash32) string {
	f, err := fsutil.Open(path)
	if err != nil {
		f.Close()
		return ""
	}
	defer f.Close()
	_, _ = io.Copy(hs, f)
	return hex.EncodeToString(hs.Sum(nil))
}

func hash64el(path string, hs hash.Hash64) string {
	f, err := fsutil.Open(path)
	if err != nil {
		f.Close()
		return ""
	}
	defer f.Close()
	_, _ = io.Copy(hs, f)
	return hex.EncodeToString(hs.Sum(nil))
}
