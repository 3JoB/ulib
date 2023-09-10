// litefmt is a simple replacement for fmt.Sprint() and fmt.Sprintln().
// It only supports string type.

package litefmt_test

import (
	"testing"

	"github.com/3JoB/go-reflect"

	"github.com/3JoB/ulib/litefmt"
)

func TestSprint(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := litefmt.Sprint(tt.args.s...); got != tt.want {
				t.Errorf("Sprint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSprintln(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := litefmt.Sprintln(tt.args.s...); got != tt.want {
				t.Errorf("Sprintln() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSprintf(t *testing.T) {
	type args struct {
		format string
		a      []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := litefmt.Sprintf(tt.args.format, tt.args.a...); got != tt.want {
				t.Errorf("Sprintf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSprintf(t *testing.T) {
	type args struct {
		format string
		a      []any
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := litefmt.BSprintf(tt.args.format, tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BSprintf() = %v, want %v", got, tt.want)
			}
		})
	}
}
