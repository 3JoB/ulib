// The GIF package provides encoding/decoding and other functions.
// It is currently a beta version and availability is not guaranteed.
package gif

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/png"

	"github.com/3JoB/vfs/memfs"

	"github.com/3JoB/ulib/fsutil"
)

type Info struct {
	// The successive delay times, one per frame, in 100ths of a second.
	Delay []int

	// FileList returns a list of all available files.
	FileList []string

	// FS returns an operational virtual file system.
	FS *memfs.MemFS
}

// Extract a GIF file.
func Decode(v string) (*Info, error) {
	if f, err := fsutil.Open(v); err != nil {
		return nil, err
	} else {
		info := &Info{
			FS: memfs.Create(),
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
			info.FS.Remove("/tmp/backend")
			return info, nil
		}
	}
}

// Create a GIF file.
func Encode(i *Info, path string) error {
	if len(i.FileList) == 0 {
		return errors.New("no files to encode")
	}
	if len(i.Delay) == 0 {
		return errors.New("the playback speed cannot be empty")
	}
	if len(i.Delay) != len(i.FileList) {
		return errors.New("the number of files does not match the playback speed")
	}
	g := &gif.GIF{}
	for e, filename := range i.FileList {
		f, err := fsutil.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		img, err := png.Decode(f)
		if err != nil {
			return err
		}
		paletted := image.NewPaletted(img.Bounds(), nil)
		paletted.Palette = append(paletted.Palette, img.At(0, 0))
		for x := 0; x < paletted.Rect.Max.X; x++ {
			for y := 0; y < paletted.Rect.Max.Y; y++ {
				paletted.Set(x, y, img.At(x, y))
			}
		}
		g.Image = append(g.Image, paletted)
		g.Delay = append(g.Delay, i.Delay[e])
	}
	out, err := fsutil.Open(path)
	if err != nil {
		return err
	}
	defer out.Close()

	if err := gif.EncodeAll(out, g); err != nil {
		return err
	}
	return nil
}
