package fsutil

import (
	"bufio"
	"os"

	"github.com/3JoB/telebot/pkg"
)

type fsutil_struct struct {
	Path  string
	Data  string
	TRUNC bool
}

func NewFile(path string) *fsutil_struct {
	fs := &fsutil_struct{
		Path: path,
	}
	return fs
}

func (f *fsutil_struct) SetTrunc() *fsutil_struct {
	if f.TRUNC {
		f.TRUNC = false
	} else {
		f.TRUNC = true
	}
	return f
}

func (f *fsutil_struct) Write(d string) error {
	var (
		file *os.File
		err  error
	)
	if f.TRUNC {
		file, err = os.OpenFile(f.Path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	} else {
		file, err = os.OpenFile(f.Path, os.O_WRONLY|os.O_CREATE, 0666)
	}
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write(pkg.Bytes(d))
	writer.Flush()
	return nil
}

type f_info struct{}

func (f *fsutil_struct) Info(filepath string) *f_info {
	return nil
}
