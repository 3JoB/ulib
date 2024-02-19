package compress

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/klauspost/compress/zip"

	"github.com/3JoB/ulib/fsutil"
)

var ErrTargetType error = errors.New("The target directory type is file")

// Extract files
func ExtractAndWriteFile(destination string, f *zip.File) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	path := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(path, filepath.Clean(destination)+string(os.PathSeparator)) {
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
		dir, _ := filepath.Split(path)
		dir = filepath.Clean(dir)
		if err := fsutil.Mkdir(dir, f.Mode()); err != nil {
			return err
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
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
