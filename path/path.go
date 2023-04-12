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
	return walker.Walk(root, fn, opts...)
}

// WalkWithContext walks the file tree rooted at root,
// calling walkFn for each file or directory in the tree, including root.
//
// If fastWalk returns filepath.SkipDir, the directory is skipped.
//
// Multiple goroutines stat the filesystem concurrently.
// The provided walkFn must be safe for concurrent use.
func WalkWithContext(ctx context.Context, root string, fn WalkFunc, opts ...walker.Option) error {
	return walker.WalkWithContext(ctx, root, fn, opts...)
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

/*
Split splits path immediately following the final Separator,
separating it into a directory and file name component.
If there is no Separator in path, Split returns an empty dir
and file set to path.

The returned values have the property that path = dir+file.
*/
func Split(src string) (dir string, file string) {
	return filepath.Split(src)
}

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is no dot.
func Ext(src string) string {
	return filepath.Ext(src)
}

/*
Dir returns all but the last element of path,
typically the path's directory. After dropping the final element,
Dir calls Clean on the path and trailing slashes are removed.
If the path is empty, Dir returns ".". If the path consists entirely
of separators, Dir returns a single separator. The returned path does
not end in a separator unless it is the root directory.
*/
func Dir(path string) string {
	return filepath.Dir(path)
}

func DirPath(src string) string {
	dir, _ := Split(src)
	return Clean(dir)
}
