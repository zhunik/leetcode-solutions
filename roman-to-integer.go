package solutions

func RomanToInt(s string) int {
	dictionary := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var value , prev, tmp int

	for i := len(s) - 1; i >= 0; i-- {
		if i == len(s)-1 {
			value = dictionary[s[i:i+1]]
			prev = dictionary[s[i:i+1]]
		} else {
			tmp = dictionary[s[i:i+1]]
			if tmp < prev {
				value -= tmp
			} else {
				value += tmp
			}
			prev = tmp
		}
	}

	return value
}
