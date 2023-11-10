package fsutil

import (
	"bufio"
	"io"
	"io/fs"
	"os"

	"github.com/3JoB/unsafeConvert"
)

const (
	O_ModePerm fs.FileMode = 511

	// O_TRUNC = os.O_WRONLY|os.O_TRUNC|os.O_CREATE
	O_TRUNC int = 577

	// O_RDWTRUNC = os.O_RDWR|os.O_CREATE|os.O_TRUNC
	O_RDWTRUNC int = 578

	// O_RDONLY = os.O_RDONLY
	O_RDONLY int = 0

	// O_WROC = os.O_WRONLY|os.O_CREATE
	O_WROC int = 65
)

func Create(v string) (*os.File, error) {
	return os.OpenFile(v, O_RDWTRUNC, 0666)
}

/*
Open opens the named file for reading.

If successful, methods on the returned file can be used for reading;
the associated file descriptor has mode O_RDONLY.

If there is an error, it will be of type *PathError.
*/
func Open(v string) (*os.File, error) {
	return os.OpenFile(v, O_RDONLY, 0)
}

func OpenRead(v string) ([]byte, error) {
	o, err := Open(v)
	if err != nil {
		return nil, err
	}
	defer o.Close()
	return io.ReadAll(o)
}

func CleanFile(path string) error {
	return TruncWrite(path, "")
}

func ReaderWriter(w io.Reader, r io.Writer) (*bufio.Reader, *bufio.Writer) {
	return bufio.NewReader(w), bufio.NewWriter(r)
}

func TruncWrite(path string, d any) error {
	file, err := os.OpenFile(path, O_TRUNC, 0666)
	if err != nil {
		file.Close()
		return err
	}
	defer file.Close()
	return write(file, d)
}

func Write(path string, d any) error {
	file, err := os.OpenFile(path, O_WROC, 0666)
	if err != nil {
		file.Close()
		return err
	}
	defer file.Close()
	return write(file, d)
}

func write(file *os.File, d any) error {
	writer := bufio.NewWriter(file)
	switch d := d.(type) {
	case string:
		writer.Write(unsafeConvert.BytePointer(d))
	case []byte:
		writer.Write(d)
	default:
		writer.Write(unsafeConvert.BytePointer(d.(string)))
	}

	return writer.Flush()
}

func Mkdir(path string, mode ...fs.FileMode) error {
	if len(mode) != 0 {
		return os.MkdirAll(path, mode[0])
	}
	return os.MkdirAll(path, O_ModePerm)
}

func Remove(v string) error {
	return os.RemoveAll(v)
}
