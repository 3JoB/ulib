// litefmt is a simple replacement for fmt.Sprint() and fmt.Sprintln().
// It only supports string type.

package litefmt

import (
	"fmt"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/strings"
)

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
	return unsafeConvert.BytesReflect(b.String())
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
	return unsafeConvert.BytesReflect(b.String())
}

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func BSprintf(format string, a ...any) []byte {
	return unsafeConvert.BytesReflect(fmt.Sprintf(format, a...))
}
