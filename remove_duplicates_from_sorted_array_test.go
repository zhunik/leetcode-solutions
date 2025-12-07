package solutions

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	cases := map[string]struct {
		nums      []int
		expNums   []int
		expLength int
	}{
		"case1": {
			nums:      []int{1, 1, 1, 2, 2, 3},
			expNums:   []int{1, 1, 2, 2, 3, 0},
			expLength: 5,
		},
		"case2": {
			nums:      []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expNums:   []int{0, 0, 1, 1, 2, 3, 3, 0, 0},
			expLength: 7,
		},
		"case3": {
			nums:      []int{1, 1},
			expNums:   []int{1, 1},
			expLength: 2,
		},
	}

	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			k := removeDuplicates(tc.nums)
			if k != tc.expLength {
				t.Errorf("expected length should be %d got %d", tc.expLength, k)
			}
			for i := 0; i < k; i++ {
				if tc.nums[i] != tc.expNums[i] {
					t.Errorf("expected %v got %v", tc.expNums, tc.nums)
					break
				}
			}
		})
	}
}
