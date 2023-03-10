package fsutil

import (
	"bufio"
	"io"
	"io/fs"
	"os"

	"github.com/3JoB/unsafeConvert"
)

func Create(v string) (*os.File, error) {
	return os.Create(v)
}

/*
CreateTemp creates a new temporary file in the directory dir,
opens the file for reading and writing, and returns the resulting file.

The filename is generated by taking pattern and adding a random string to the end.

If pattern includes a "*", the random string replaces the last "*".

If dir is the empty string, CreateTemp uses the default directory for temporary files, as returned by TempDir.

Multiple programs or goroutines calling CreateTemp simultaneously will not choose the same file.

The caller can use the file's Name method to find the pathname of the file.

It is the caller's responsibility to remove the file when it is no longer needed.
*/
func CreateTemp(dir string, pattern string) (*os.File, error) {
	return os.CreateTemp(dir, pattern)
}

/*
Open opens the named file for reading.

If successful, methods on the returned file can be used for reading;
the associated file descriptor has mode O_RDONLY.

If there is an error, it will be of type *PathError.
*/
func Open(v string) (*os.File, error) {
	return os.Open(v)
}

func OpenRead(v string) ([]byte, error) {
	o, err := Open(v)
	if err != nil {
		return nil, err
	}
	defer o.Close()
	if data, err := ReadAll(o); err != nil {
		return nil, err
	} else {
		return data, err
	}
}

/*
ReadAll reads from r until an error or EOF and returns the data it read.
A successful call returns err == nil, not err == EOF.
Because ReadAll is defined to read from src until EOF,
it does not treat an EOF from Read as an error to be reported.
*/
func ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

/*
OpenFile is the generalized open call;
most users will use Open or Create instead.
It opens the named file with specified flag (O_RDONLY etc.).
If the file does not exist, and the O_CREATE flag is passed,
it is created with mode perm (before umask).
If successful, methods on the returned File can be used for I/O.
If there is an error, it will be of type *PathError.
*/
func OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func TruncWrite(path string, d any) error {
	file, err := OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		file.Close()
		return err
	}
	defer file.Close()
	return write(file, d)
}

func Write(path string, d any) error {
	file, err := OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
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
		writer.Write(unsafeConvert.Bytes(d))
	case []byte:
		writer.Write(d)
	default:
		writer.Write(unsafeConvert.Bytes(d.(string)))
	}
	writer.Flush()
	return nil
}

func Mkdir(path string, mode ...fs.FileMode) error {
	if len(mode) != 0 {
		return os.MkdirAll(path, mode[0])
	}
	return os.MkdirAll(path, os.ModePerm)
}

func Remove(v string) error {
	return os.RemoveAll(v)
}
