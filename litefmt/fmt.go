// litefmt is a simple replacement for fmt.Sprint() and fmt.Sprintln().
// It only supports string type.

package litefmt

import (
	"fmt"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/strings"
)

var builders = strings.NewBuilders()

func LSprint(s ...string) string {
	for _, r := range s {
		builders.WriteString(r)
	}
	defer builders.Reset()
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
