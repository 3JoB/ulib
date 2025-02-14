package zszip

import (
	"io"
	"os"

	"github.com/klauspost/compress/zip"
	zs "github.com/klauspost/compress/zstd"

	"github.com/3JoB/ulib/fsutil"
	"github.com/3JoB/ulib/fsutil/compress"
)

type Zip struct{}

func New() *Zip {
	return &Zip{}
}

// Create compresses the specified files into a zip archive at the given source path. Existing files at the source path
// are overwritten. It uses specific compression methods while processing files. Returns an error if any step fails.
func (z Zip) Create(source string, files []string) error {
	if fsutil.IsExist(source) {
		fsutil.Remove(source)
	}
	fs, err := os.OpenFile(source, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	encomp := zs.ZipCompressor(zs.WithEncoderLevel(zs.EncoderLevelFromZstd(17)))
	w := zip.NewWriter(fs)
	w.RegisterCompressor(zs.ZipMethodPKWare, encomp)
	w.RegisterCompressor(zs.ZipMethodWinZip, encomp)
	for _, f := range files {
		ofs, err := os.OpenFile(f, os.O_RDWR, 0755)
		if err != nil {
			return err
		}
		zfs, err := w.CreateHeader(&zip.FileHeader{
			Name:   f,
			Method: zs.ZipMethodWinZip,
		})
		if err != nil {
			return err
		}
		if _, err := io.Copy(zfs, ofs); err != nil {
			return err
		}
		ofs.Close()
	}
	w.Close()
	return nil
}

// Extract extracts files from a zip archive at the specified source path to the given destination directory.
// It returns a slice of extracted file names and an error if the operation fails.
func (z Zip) Extract(source, destination string) (extractedFiles []string, err error) {
	decomp := zs.ZipDecompressor()
	zip.RegisterDecompressor(zs.ZipMethodPKWare, decomp)
	zip.RegisterDecompressor(zs.ZipMethodWinZip, decomp)
	r, err := zip.OpenReader(source)
	if err != nil {
		return nil, err
	}

	defer r.Close()

	if !fsutil.IsExist(destination) {
		if err := fsutil.Mkdir(destination, 0755); err != nil {
			return nil, err
		}
	} else {
		if !fsutil.IsDir(destination) {
			return nil, compress.ErrTargetType
		}
	}

	for _, f := range r.File {
		if err := compress.ExtractAndWriteFile(destination, f); err != nil {
			return nil, err
		}

		extractedFiles = append(extractedFiles, f.Name)
	}

	return extractedFiles, nil
}
