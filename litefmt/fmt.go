// litefmt is a simple replacement for fmt.Sprint() and fmt.Sprintln().
// It only supports string type.

package litefmt

import (
	"bytes"
	"fmt"
	strs "strings"
	"sync"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/strings"
)

var builders = strs.Builder{}

var sprintPool = sync.Pool{
	New: func() any {
		return &strs.Builder{}
	},
}

func LSprint(s ...string) string {
	for _, r := range s {
		builders.WriteString(r)
	}
	defer builders.Reset()
	return builders.String()
}

func VSprint(s ...string) string {
	bs := bytes.Buffer{}
	bs.Grow(len(s))
	for _, r := range s {
		bs.WriteString(r)
	}
	bs.Reset()
	return bs.String()
}

func PSprint(s ...string) string {
	b := sprintPool.Get().(*strs.Builder)
	for _, r := range s {
		b.WriteString(r)
	}
	defer b.Reset()
	defer sprintPool.Put(b)
	return builders.String()
}

func TSprint(s ...string) string {
	b := strings.NewBuilders()
	var buf []byte
	for _, r := range s {
		buf = append(buf, r...)
	}
	b.Write(buf)
	return b.String()
}

func Sprint(s ...string) string {
	b := strings.NewBuilders()
	for _, r := range s {
		b.WriteString(r)
	}
	return b.String()
}

func BSprint(s ...string) []byte {
	b := strings.NewBuilders()
	for _, r := range s {
		b.WriteString(r)
	}
	return unsafeConvert.BytePointer(b.String())
}

func Sprintln(s ...string) string {
	b := strings.NewBuilders()
	for _, r := range s {
		b.WriteString(r)
	}
	b.WriteString("\n")
	return b.String()
}

func BSprintln(s ...string) []byte {
	b := strings.NewBuilders()
	for _, r := range s {
		b.WriteString(r)
	}
	b.WriteString("\n")
	return unsafeConvert.BytePointer(b.String())
}

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func BSprintf(format string, a ...any) []byte {
	return unsafeConvert.BytePointer(fmt.Sprintf(format, a...))
}
