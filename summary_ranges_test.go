package solutions

import "testing"

func TestSummaryRanges(t *testing.T) {
	cases := map[string]struct {
		nums   []int
		expRes []string
	}{
		"case1": {
			nums:   []int{0, 1, 2, 4, 5, 7},
			expRes: []string{"0->2", "4->5", "7"},
		},
		"case2": {
			nums:   []int{0, 2, 3, 4, 6, 8, 9},
			expRes: []string{"0", "2->4", "6", "8->9"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res := summaryRanges(tc.nums)
			if len(res) != len(tc.expRes) {
				t.Errorf("expected result %v, got %v", tc.expRes, res)
				t.Fail()
			}
			for k, v := range res {
				if tc.expRes[k] != v {
					t.Errorf("expected range is %q, got %q ", tc.expRes[k], v)
				}
			}
		})
	}
}
