package solutions

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	cases := map[string]struct{
		strs []string
		expPref string
	}{
		"case1":{
			strs: []string{},
			expPref: "",
		},
		"case2":{
			strs: []string{"flower","flow","flight"},
			expPref: "fl",
		},
		"case3": {
			strs: []string{"dog","racecar","car"},
			expPref: "",
		},
		"case4": {
			strs: []string{"ab","a"},
			expPref: "a",
		},

	}

	for n,tc := range cases {
		t.Run(n, func(t *testing.T) {
			r := longestCommonPrefix(tc.strs)
			if r != tc.expPref {
				t.Errorf("expected prefix %q got %q", tc.expPref, r)
			}
		})
	}
}