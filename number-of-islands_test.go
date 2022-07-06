package solutions

import "testing"

func TestNumIslands(t *testing.T) {
	cases := []struct {
		grid [][]byte
		want int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
	}

	for _, tc := range cases {
		tc := tc
		got := numIslands(tc.grid)
		if got != tc.want {
			t.Errorf("Wrong number of islands in %+v want %d got %d", tc.grid, tc.want, got)
		}
	}

}
