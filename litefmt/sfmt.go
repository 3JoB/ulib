// litefmt is a simple replacement for fmt.Sprint() and fmt.Sprintln().
// It only supports string type.

package litefmt

import (
	"fmt"
	"sync"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/strings"
)

var (
	bytesPool sync.Pool
)

// The most primitive Sprint implementation, only allows string values to pass.
func Sprint(s ...string) string {
	b := strings.NewBuilders()
	for _, r := range s {
		b.WriteString(r)
	}
	return b.String()
}

// The pooled Sprint method. Of course, you don't need to deal with it.
// Although pointers will be generated, they are automatically managed by the pool.
//
// Oh, right. This PSprint is re-implemented, and its performance far exceeds the old PSprint.
func PSprint(s ...string) string {
	b := psp_acquire()
	for _, r := range s {
		b.WriteString(r)
	}
	defer psp_release(b)
	return b.String()
}

// The pooled Sprint method. Of course, you don't need to deal with it.
// Although pointers will be generated, they are automatically managed by the pool.
//
// Oh, right. This PSprint is re-implemented, and its performance far exceeds the old PSprint.
//
// The difference from PSprint is that it uses unsafe to convert []byte to string.
// Although it returns an immutable value, it can get a huge performance improvement.
func PSprintP(s ...string) string {
	b := psp_acquire()
	for _, r := range s {
		b.WriteString(r)
	}
	defer psp_release(b)

	return unsafeConvert.StringSlice(b.Bytes())
}

func Sprintln(s ...string) string {
	b := psp_acquire()
	for _, r := range s {
		b.WriteString(r)
	}
	b.WriteString("\n")
	defer psp_release(b)
	return b.String()
}

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func BSprintf(format string, a ...any) []byte {
	return unsafeConvert.BytePointer(fmt.Sprintf(format, a...))
}
