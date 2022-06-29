package solutions

import (
	"fmt"
	"testing"
)

// ["LRUCache","put","put","get","put","get","put","get","get","get"]
// [[2},{1,1},{2,2},{1},{3,3},{2},{4,4},{1},{3},{4]]
//
func runTests(commands []string, args [][]int, expected []int, t *testing.T) {
	var lru LRUCache
	var got []int

	fmt.Println(fmt.Sprintf("%+v", got))

	for i := 0; i < len(commands); i++ {
		command := commands[i]

		switch command {
		case "LRUCache":
			lru = Constructor(args[i][0])
			got = append(got, 0)
		case "put":
			lru.Put(args[i][0], args[i][1])
			got = append(got, 0)
		case "get":
			res := lru.Get(args[i][0])
			got = append(got, res)
		default:
			fmt.Println("default command")
		}

		fmt.Println(fmt.Sprintf("running %s %+v, expected %d got %d", command, args[i], expected[i], got[i]))

		if expected[i] != got[i] {
			t.Error(fmt.Sprintf("Command %s %+v failed: expected %d, got %d", command, args[i], expected[i], got[i]))
		}
	}
}

type testCase struct {
	commands []string
	args     [][]int
	expected []int
}

func TestLRUCache(t *testing.T) {
	cases := []testCase{
		{
			commands: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			args:     [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			expected: []int{0, 0, 0, 1, 0, -1, 0, -1, 3, 4},
		},
		{
			commands: []string{"LRUCache", "put", "put", "get", "put", "put", "get"},
			args:     [][]int{{2}, {2, 1}, {2, 2}, {2}, {1, 1}, {4, 1}, {2}},
			expected: []int{0, 0, 0, 2, 0, 0, -1},
		},
	}

	for _, tc := range cases {
		runTests(tc.commands, tc.args, tc.expected, t)
	}

}
