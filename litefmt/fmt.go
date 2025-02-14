package litefmt

import (
	"github.com/3JoB/ulib/pool"
	"os"
)

func Print(s ...string) {
	b := pool.NewBuffer()
	for _, r := range s {
		b.WriteString(r)
	}
	defer pool.ReleaseBuffer(b)
	// io.Discard.Write(b.Bytes())
	os.Stdout.Write(b.Bytes())
}

func Println(s ...string) {
	b := pool.NewBuffer()
	for _, r := range s {
		b.WriteString(r)
	}
	b.WriteString("\n")
	defer pool.ReleaseBuffer(b)
	os.Stdout.Write(b.Bytes())
}
