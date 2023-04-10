package path

import (
	"context"
	"os"
	_ "unsafe"

	"github.com/saracen/walker"
)

type WalkFunc func(pathname string, fi os.FileInfo) error

type WalkCallback func(pathname string, err error) error

// Walk wraps WalkWithContext using the background context.
//
//go:linkname Walk github.com/saracen/walker.Walk
func Walk(root string, fn WalkFunc, opts ...walker.Option) error

// WalkWithContext walks the file tree rooted at root,
// calling walkFn for each file or directory in the tree, including root.
//
// If fastWalk returns filepath.SkipDir, the directory is skipped.
//
// Multiple goroutines stat the filesystem concurrently.
// The provided walkFn must be safe for concurrent use.
//
//go:linkname WalkWithContext github.com/saracen/walker.WalkWithContext
func WalkWithContext(ctx context.Context, root string, fn WalkFunc, opts ...walker.Option) error

// WithErrorCallback sets a callback to be used for error handling.
// Any error returned will halt the Walk function and return the error.
// If the callback returns nil Walk will continue.
//
//go:linkname WithErrorCallback github.com/saracen/walker.WithErrorCallback
func WithErrorCallback(call WalkCallback) walker.Option

//go:linkname Abs filepath.Abs
func Abs(v string) (string, error)

//go:linkname IsAbs filepath.IsAbs
func IsAbs(path string) bool

//go:linkname IsLocal filepath.IsLocal
func IsLocal(path string) bool

//go:linkname Join filepath.Join
func Join(v ...string) string

//go:linkname Base filepath.Base
func Base(src string) string

//go:linkname Clean filepath.Clean
func Clean(src string) string

/*Split splits path immediately following the final Separator,
separating it into a directory and file name component.
If there is no Separator in path, Split returns an empty dir
and file set to path.

The returned values have the property that path = dir+file.
*/
//go:linkname Split filepath.Split
func Split(src string) (dir string, file string)

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is no dot.
//
//go:linkname Ext filepath.Ext
func Ext(src string) string

/*Dir returns all but the last element of path,
typically the path's directory. After dropping the final element,
Dir calls Clean on the path and trailing slashes are removed.
If the path is empty, Dir returns ".". If the path consists entirely
of separators, Dir returns a single separator. The returned path does
not end in a separator unless it is the root directory.
*/
//go:linkname Dir filepath.Dir
func Dir(path string) string

func DirPath(src string) string {
	dir, _ := Split(src)
	return Clean(dir)
}
