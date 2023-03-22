# ULIB Keyword

Fast wrapper for keyword lookup.

Example: [Github](https://github.com/3JoB/ulib/blob/master/test/keyword/main.go)

## Regexp

Normal
```go
package main

import (
	"fmt"

	"github.com/3JoB/ulib/keyword"
)

func main() {
	text := "hello world"
	kw := "world"
	fmt.Printf("RegexpFind: %v\n", keyword.RegexpFind(text, kw))
}
```

Compile:
```go
package main

import (
    "fmt"

    "github.com/3JoB/ulib/keyword"
)

func main(){
    text := "hello world"
    text2 := "hello golang"
	kw := "world"
    re, _ := keyword.RegexpCompile(kw)
	reb, _ := re.Find(text)
	fmt.Printf("RegexpFind: %v\n", reb)
    reb2, _ := re.Find(text2)
	fmt.Printf("RegexpFind 2: %v\n", reb2)
}
```