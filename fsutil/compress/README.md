## compress

form: https://github.com/artdarek/go-unzip

Wrapper for Zip and 7zip decompression.


### Example

```go
package main

import (
	"fmt"

	"github.com/ulib/fsutil/compress"
)

func main() {
	u := compress.NewSevenZip()

	files, err := u.Extract("./data/file.zip", "./data/directory")
	if err != nil {
		panic(err)
	}

	fmt.Printf("extracted files count: %d \n", len(files))
	fmt.Printf("files list: %v \n", files)
}
```