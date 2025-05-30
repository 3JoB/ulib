package rands

import (
	"testing"
)

func TestRand(t *testing.T) {
	tests := []struct {
		name      string
		algorithm int
		input     []int
		num       int
		want      int // expected length of result
	}{
		{name: "stdV1_positive", algorithm: STD, input: []int{1, 2, 3, 4, 5}, num: 3, want: 3},
		{name: "stdV1_empty_slice", algorithm: STD, input: []int{}, num: 3, want: 0},
		{name: "stdV1_zero_num", algorithm: STD, input: []int{1, 2, 3, 4, 5}, num: 0, want: 0},
		{name: "stdV1_num_exceeds_input", algorithm: STD, input: []int{1, 2, 3}, num: 5, want: 0},
		{name: "stdV2_positive", algorithm: STDV2, input: []int{1, 2, 3, 4, 5}, num: 3, want: 3},
		{name: "stdV2_empty_slice", algorithm: STDV2, input: []int{}, num: 3, want: 0},
		{name: "stdV2_zero_num", algorithm: STDV2, input: []int{1, 2, 3, 4, 5}, num: 0, want: 0},
		{name: "stdV2_num_exceeds_input", algorithm: STDV2, input: []int{1, 2, 3}, num: 5, want: 0},
		{name: "frand_positive", algorithm: FRAND, input: []int{1, 2, 3, 4, 5}, num: 3, want: 3},
		{name: "frand_empty_slice", algorithm: FRAND, input: []int{}, num: 3, want: 0},
		{name: "frand_zero_num", algorithm: FRAND, input: []int{1, 2, 3, 4, 5}, num: 0, want: 0},
		{name: "frand_num_exceeds_input", algorithm: FRAND, input: []int{1, 2, 3}, num: 5, want: 0},
		{name: "pg_positive", algorithm: PG, input: []int{1, 2, 3, 4, 5}, num: 3, want: 3},
		{name: "pg_empty_slice", algorithm: PG, input: []int{}, num: 3, want: 0},
		{name: "pg_zero_num", algorithm: PG, input: []int{1, 2, 3, 4, 5}, num: 0, want: 0},
		{name: "pg_num_exceeds_input", algorithm: PG, input: []int{1, 2, 3}, num: 5, want: 0},
		{name: "unknown_algorithm", algorithm: 999, input: []int{1, 2, 3, 4, 5}, num: 3, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rand(tt.algorithm, tt.input, tt.num)
			if got == nil && tt.want != 0 {
				t.Errorf("expected length %d, but got nil", tt.want)
			}
			if got != nil && len(got) != tt.want {
				t.Errorf("expected length %d, but got %d", tt.want, len(got))
			}
		})
	}
}
