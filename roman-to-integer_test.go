package solutions

import "testing"

func TestRomanToInt(t *testing.T) {

	results := map[string]int {
		"III": 3,
		"IX": 9,
		"XI": 11,
		"LVIII": 58,
		"MCMXCIV": 1994,
	}

	for roman, value := range results {
		got := RomanToInt(roman)
		if got != value {
			t.Errorf("got '%d' want '%d'", got, value)
		}
	}

}
