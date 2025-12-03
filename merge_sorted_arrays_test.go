package solutions

import "testing"

func TestMergeSortedArrays(t *testing.T) {

	cases := map[string]struct{
		nums1 []int
		m int
		nums2 []int
		n int
		expected []int
	}{
		"case1": {
			nums1: []int{1,2,3,0,0,0}, m: 3,
			nums2: []int{2,5,6}, n : 3,
			expected: []int{1,2,2,3,5,6},
		},
		"case2": {
			nums1: []int{1}, m: 1,
			nums2: []int{}, n : 0,
			expected: []int{1},
		},

		"case3": {
			nums1: []int{0}, m: 0,
			nums2: []int{1}, n : 1,
			expected: []int{1},
		},
		"case4": {
			nums1: []int{4,0,0,0,0,0}, m: 1,
			nums2: []int{1,2,3,5,6}, n : 5,
			expected: []int{1,2,3,4,5,6},
		},
		"case5": {
			nums1: []int{1,2,4,5,6,0}, m: 5,
			nums2: []int{3}, n : 1,
			expected: []int{1,2,3,4,5,6},
		},

	}

	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			mergeSortedArrays(tc.nums1, tc.m, tc.nums2, tc.n)
			for k,v := range tc.expected {
				if tc.nums1[k] != v {
					t.Errorf("expected array should be %v got %v", tc.expected, tc.nums1)
					break
				}
			}
		})
	}
}