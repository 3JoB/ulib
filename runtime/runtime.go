package runtime

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// IsInTest determines if the current code is running within a test environment by inspecting the runtime function name.
func IsInTest() bool {
	return strings.HasPrefix(runtime.FuncForPC(reflect.ValueOf(testing.RunTests).Pointer()).Name(), "testing.RunTests")
}
