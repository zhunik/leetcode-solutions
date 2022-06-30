package solutions

import (
	"fmt"
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestAsteroidCollision(t *testing.T) {
	type testCase struct {
		input []int
		want  []int
	}
	cases := []testCase{
		{
			input: []int{5, 10, -5},
			want:  []int{5, 10},
		},
		{
			input: []int{8, -8},
			want:  []int{},
		},
		{
			input: []int{10, 2, -5},
			want:  []int{10},
		},
		{
			input: []int{-2, -1, 1, 2},
			want:  []int{-2, -1, 1, 2},
		},
	}

	for _, tc := range cases {
		got := asteroidCollision(tc.input)
		if !equal(got, tc.want) {
			t.Error(fmt.Sprintf("wrong collision %+v wanted %+v got %+v", tc.input, tc.want, got))
		}
	}
}
