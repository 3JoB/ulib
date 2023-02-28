package compress

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/bodgit/sevenzip"
)

type sevenZip struct {
	pass string
}

func New7Zip(pwd ...string) *sevenZip {
	if len(pwd) != 0 {
		return &sevenZip{pass: pwd[0]}
	}
	return &sevenZip{}
}

func (sv sevenZip) Extract(source, destination string) (extractedFiles []string, err error) {
	var i *sevenzip.ReadCloser
	if sv.pass != "" {
		i, err = sevenzip.OpenReaderWithPassword(source, sv.pass)
	} else {
		i, err = sevenzip.OpenReader(source)
	}
	if err != nil {
		return nil, err
	}
	defer i.Close()
	if err := os.MkdirAll(destination, 0755); err != nil {
		return nil, err
	}

	for _, f := range i.File {
		if err := sv.extractAndWriteFile(destination, f); err != nil {
			return nil, err
		}

		extractedFiles = append(extractedFiles, f.Name)
	}

	return extractedFiles, nil
}

func (sevenZip) extractAndWriteFile(destination string, f *sevenzip.File) error {
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
		if err = os.MkdirAll(path, f.Mode()); err != nil {
			return err
		}
	} else {
		err = os.MkdirAll(filepath.Dir(path), f.Mode())
		if err != nil {
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
