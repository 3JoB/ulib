package iocopy

import (
	"io"
	"os"
	"syscall"
)

func Copy(dst io.Writer, src io.Reader) (written int64, err error) {
    hDst := syscall.Handle(dst.(*os.File).Fd())
    hSrc := syscall.Handle(src.(*os.File).Fd())
    var n int
    if err = asmCopy(hSrc, hDst, &n); err != nil {
        return 0, err
    }
    return int64(n), nil
} 

func asmCopy(hSrc, hDst syscall.Handle, n *int) error