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
		{"stdV1_positive", STD, []int{1, 2, 3, 4, 5}, 3, 3},
		{"stdV1_empty_slice", STD, []int{}, 3, 0},
		{"stdV1_zero_num", STD, []int{1, 2, 3, 4, 5}, 0, 0},
		{"stdV1_num_exceeds_input", STD, []int{1, 2, 3}, 5, 0},
		{"stdV2_positive", STDV2, []int{1, 2, 3, 4, 5}, 3, 3},
		{"stdV2_empty_slice", STDV2, []int{}, 3, 0},
		{"stdV2_zero_num", STDV2, []int{1, 2, 3, 4, 5}, 0, 0},
		{"stdV2_num_exceeds_input", STDV2, []int{1, 2, 3}, 5, 0},
		{"frand_positive", FRAND, []int{1, 2, 3, 4, 5}, 3, 3},
		{"frand_empty_slice", FRAND, []int{}, 3, 0},
		{"frand_zero_num", FRAND, []int{1, 2, 3, 4, 5}, 0, 0},
		{"frand_num_exceeds_input", FRAND, []int{1, 2, 3}, 5, 0},
		{"pg_positive", PG, []int{1, 2, 3, 4, 5}, 3, 3},
		{"pg_empty_slice", PG, []int{}, 3, 0},
		{"pg_zero_num", PG, []int{1, 2, 3, 4, 5}, 0, 0},
		{"pg_num_exceeds_input", PG, []int{1, 2, 3}, 5, 0},
		{"unknown_algorithm", 999, []int{1, 2, 3, 4, 5}, 3, 0},
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
