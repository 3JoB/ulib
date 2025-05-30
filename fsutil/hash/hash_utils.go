package hash

import (
	"fmt"
	"hash"
	"io"

	"github.com/3JoB/ulib/fsutil"
)

func hashel(path string, hs hash.Hash) string {
	f, err := fsutil.Open(path)
	if err != nil {
		f.Close()
		return ""
	}
	defer f.Close()
	_, _ = io.Copy(hs, f)
	return fmt.Sprintf("%x", hs.Sum(nil))
}

func hashrl(r io.Reader, hs hash.Hash) string {
	defer func() {
		hs.Reset()
		if s, ok := r.(io.Seeker); ok {
			_, _ = s.Seek(0, io.SeekStart)
		}
	}()
	if _, err := io.Copy(hs, r); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", hs.Sum(nil))
}
