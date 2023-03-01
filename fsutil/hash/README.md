# hash

This package is used to calculate file hash.

The `NewWithOs` method will automatically call the `os.File.Close()` method by default. To prevent the call, see the following example.

```go
package main

import (
    "fmt"
    "os"

	"github.com/3JoB/ulib/fsutil/hash"
)

func main(){
    y, err := os.Open("GMakefile.yml")
    if err != nil {
        return nil
    }
    md5, _ := hash.NewWithOs(y).DisableAutoClose().MD5()
    fmt.Println(md5)
}
```