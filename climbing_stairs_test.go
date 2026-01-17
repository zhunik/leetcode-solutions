package solutions

import "testing"


func TestClimbStairs(t *testing.T) {

	cases := map[string]struct{
		n int
		r int
	}{
		"case1" : {
			n: 1,
			r: 1,
		},
		"case2" : {
			n: 2,
			r: 2,
		},
		"case3" : {
			n: 3,
			r: 3,
		},		
		"case4" : {
			n: 4,
			r: 5,
		},
		"case5" : {
			n: 5,
			r: 8,
		},		

	}


	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			r := climbStairs(tc.n)
			if r != tc.r {
				t.Errorf("expecting %d, got %d", tc.r, r)
			}
		})
	}
	
}