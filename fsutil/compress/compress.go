package compress

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/3JoB/ulib/fsutil"
	"github.com/klauspost/compress/zip"
)

var ErrTargetType = errors.New("The target directory type is file")

// ExtractAndWriteFile extracts a file from a zip archive and writes it to the specified destination path.
// If the file is a directory, it ensures the directory exists. Otherwise, writes the file content to the destination.
// Returns an error if path validation fails, if file writing fails, or if creating directories encounters issues.
func ExtractAndWriteFile(destination string, zipFile *zip.File) error {
	reader, err := zipFile.Open()
	if err != nil {
		return err
	}
	defer reader.Close()

	targetPath := filepath.Join(destination, zipFile.Name)
	if !isPathSafe(destination, targetPath) {
		return fmt.Errorf("%s: illegal file path", targetPath)
	}

	if zipFile.FileInfo().IsDir() {
		if !fsutil.IsExist(targetPath) {
			return fsutil.Mkdir(targetPath, zipFile.Mode())
		}
		return nil
	}

	if fsutil.IsExist(targetPath) {
		if err := fsutil.Remove(targetPath); err != nil {
			return err
		}
	}

	dir := filepath.Clean(filepath.Dir(targetPath))
	if err := fsutil.Mkdir(dir, zipFile.Mode()); err != nil {
		return err
	}

	destFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipFile.Mode())
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, reader); err != nil {
		return err
	}
	return nil
}

// isPathSafe checks if the targetPath is within the destination directory, ensuring no path traversal issues.
func isPathSafe(destination, targetPath string) bool {
	cleanDest := filepath.Clean(destination) + string(os.PathSeparator)
	return strings.HasPrefix(targetPath, cleanDest)
}
