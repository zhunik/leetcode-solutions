package solutions

import "testing"

func TestIntToRoman(t *testing.T) {

	results := map[int]string{
		3:    "III",
		9:    "IX",
		11:   "XI",
		58:   "LVIII",
		1994: "MCMXCIV",
	}

	for int, roman := range results {
		got := IntToRoman(int)
		if got != roman {
			t.Errorf("got '%s' want '%s'", got, roman)
		}
	}

}
