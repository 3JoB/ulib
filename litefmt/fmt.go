package litefmt

import (
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/goccy/go-reflect"
)

func Print(s ...string) {
	b := psp_acquire()
	for _, r := range s {
		b.WriteString(r)
	}
	defer psp_release(b)
	// io.Discard.Write(b.Bytes())
	os.Stdout.Write(b.Bytes())
}

func Println(s ...string) {
	b := psp_acquire()
	for _, r := range s {
		b.WriteString(r)
	}
	b.WriteString("\n")
	defer psp_release(b)
	os.Stdout.Write(b.Bytes())
}

func IsInTest() bool {
	return strings.HasPrefix(runtime.FuncForPC(reflect.ValueOf(testing.RunTests).Pointer()).Name(), "testing.RunTests")
}
