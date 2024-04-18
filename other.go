package ulib

import (
	"runtime"
	"strings"
	"testing"

	"github.com/goccy/go-reflect"
)

func IsInTest() bool {
	return strings.HasPrefix(runtime.FuncForPC(reflect.ValueOf(testing.RunTests).Pointer()).Name(), "testing.RunTests")
}
