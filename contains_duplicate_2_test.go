package solutions

import "testing"


func TestContainsNearbyDuplicate(t *testing.T) {
	cases := map[string]struct{
		nums []int
		k int
		exp bool
	}{
		"case1": {
			nums: []int{1,2,3,1},
			k: 3,
			exp: true,
		},
		"case2": {
			nums: []int{1,0,1,1},
			k: 1,
			exp: true,
		},
		"case3": {
			nums: []int{1,2,3,1,2,3},
			k: 2,
			exp: false,
		},
	}
	
	for n,tc := range cases {
		t.Run(n,func(t *testing.T) {
			res := containsNearbyDuplicate(tc.nums, tc.k)
			if res != tc.exp {
				t.Errorf("case %s expected %v got %v", n, tc.exp, res)
			}
		})
	}
}