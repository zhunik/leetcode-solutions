package solutions

import "testing"

func TestSearchMatrix(t *testing.T) {

	m1 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}

	t1 := 3

	r1 := true

	got := SearchMatrix(m1, t1)
	if got != r1 {
		t.Errorf("got '%v' want '%v'", got, r1)
	}

	m2 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}

	t2 := 13

	r2 := false

	got2 := SearchMatrix(m2, t2)

	if got2 != r2 {
		t.Errorf("got '%v' want '%v'", got2, r2)
	}

}
