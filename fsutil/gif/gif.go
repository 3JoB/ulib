// The GIF package provides encoding/decoding/super-resolution (possible) and other functions.
// It is currently a beta version and availability is not guaranteed.
package gif

import (
	"fmt"
	"image/gif"
	"image/png"

	"github.com/blang/vfs"

	"github.com/3JoB/ulib/fsutil"
)

type Info struct {
	// The successive delay times, one per frame, in 100ths of a second.
	Delay    []int
	// FileList returns a list of all available files.
	FileList []string
	// FS returns an operational virtual file system.
	FS       vfs.Filesystem
}

// Extract a GIF file.
func Decode(v string) (*Info, error) {
	if f, err := fsutil.Open(v); err != nil {
		return nil, err
	} else {
		info := &Info{
			FS: vFS(),
		}
		if gifs, err := gif.DecodeAll(f); err != nil {
			return nil, err
		} else {
			info.Delay = gifs.Delay
			if err := info.FS.Mkdir("/tmp/backend", 0755); err != nil {
				return nil, err
			}
			for i, t := range gifs.Image {
				filename := fmt.Sprintf("/tmp/backend/backend_%v.png", i)
				info.FileList = append(info.FileList, filename)
				if of, err := info.FS.OpenFile(filename, fsutil.O_RDWTRUNC, 0755); err != nil {
					return nil, err
				} else {
					png.Encode(of, t)
					of.Close()
				}
			}
			return info, nil
		}
	}
}

// Filesystem represents an abstract filesystem
func vFS() vfs.Filesystem {
	var s vfs.Filesystem
	return s
}
