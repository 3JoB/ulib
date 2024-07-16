package runtime

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func IsInTest() bool {
	return strings.HasPrefix(runtime.FuncForPC(reflect.ValueOf(testing.RunTests).Pointer()).Name(), "testing.RunTests")
}
