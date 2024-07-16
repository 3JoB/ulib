package rands_test

import (
	"reflect"
	"testing"

	"github.com/3JoB/ulib/rands"
)

func TestRands(t *testing.T) {
	type args struct {
		n   []string
		num int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rands.Rands(tt.args.n, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCRands(t *testing.T) {
	type args struct {
		n   []string
		num int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rands.CRands(tt.args.n, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CRands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandStd(t *testing.T) {
	type args struct {
		n   []string
		num int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rands.RandStd(tt.args.n, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandStd() = %v, want %v", got, tt.want)
			}
		})
	}
}
