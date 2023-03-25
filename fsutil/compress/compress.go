package compress

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/klauspost/compress/zip"

	"github.com/3JoB/ulib/err"
	"github.com/3JoB/ulib/fsutil"
	ph "github.com/3JoB/ulib/path"
)

var ErrTargetType error = &err.Err{Op: "ulib.fsutil.compress", Err: "The target directory type is file"}

// Extract files
func ExtractAndWriteFile(destination string, f *zip.File) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	path := ph.Join(destination, f.Name)
	if !strings.HasPrefix(path, ph.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("%s: illegal file path", path)
	}

	if f.FileInfo().IsDir() {
		if !fsutil.IsExist(path) {
			if err = fsutil.Mkdir(path, f.Mode()); err != nil {
				return err
			}
		}
	} else {
		if fsutil.IsExist(path) {
			if err := fsutil.Remove(path); err != nil {
				return err
			}
		}
		if err := fsutil.Mkdir(ph.DirPath(path), f.Mode()); err != nil {
			return err
		}

		f, err := fsutil.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(f, rc); err != nil {
			return err
		}
	}

	return nil
}
