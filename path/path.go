package path

import (
	"context"
	"os"
	"path/filepath"

	"github.com/saracen/walker"
)

type WalkFunc func(pathname string, fi os.FileInfo) error
type WalkCallback func(pathname string, err error) error

// Walk wraps WalkWithContext using the background context.
func Walk(root string, fn WalkFunc, opts ...walker.Option) error {
	if len(opts) != 0 {
		return walker.Walk(root, fn, opts...)
	}
	return walker.Walk(root, fn)
}

// WalkWithContext walks the file tree rooted at root, 
// calling walkFn for each file or directory in the tree, including root.
//
//If fastWalk returns filepath.SkipDir, the directory is skipped.
//
//Multiple goroutines stat the filesystem concurrently. 
// The provided walkFn must be safe for concurrent use.
func WalkWithContext(ctx context.Context, root string, fn WalkFunc, opts ...walker.Option) error {
	if len(opts) != 0 {
		return walker.WalkWithContext(ctx, root, fn, opts...)
	}
	return walker.WalkWithContext(ctx, root, fn)
}

// WithErrorCallback sets a callback to be used for error handling. 
// Any error returned will halt the Walk function and return the error. 
// If the callback returns nil Walk will continue.
func WithErrorCallback(call WalkCallback) walker.Option {
	return walker.WithErrorCallback(call)
}

func Abs(v string) (string, error) {
	return filepath.Abs(v)
}

func IsAbs(path string) bool {
	return filepath.IsAbs(path)
}

func IsLocal(path string) bool {
	return filepath.IsLocal(path)
}

func Join(v ...string) string {
	return filepath.Join(v...)
}

func Base(src string) string {
	return filepath.Base(src)
}

func Clean(src string) string {
	return filepath.Clean(src)
}

func Split(src string) (string, string) {
	return filepath.Split(src)
}

func Ext(src string) string {
	return filepath.Ext(src)
}

func Dir(path string) string {
	return filepath.Dir(path)
}

func DirPath(src string) string {
	dir, _ := Split(src)
	return Clean(dir)
}
