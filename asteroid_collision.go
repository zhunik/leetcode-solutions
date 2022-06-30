package solutions

func asteroidCollision(asteroids []int) []int {
	ans := make([]int, len(asteroids))
	l := 0
	for _, x := range asteroids {
		if x > 0 { // if x is positive just add
			ans[l] = x
			l++
		} else {
			for l > 0 && ans[l-1] > 0 && ans[l-1] < -x { // remove all positive whose value is less
				l--
			}
			if l == 0 || ans[l-1] < 0 { // is size 0 or last one is also negative
				ans[l] = x
				l++
			} else if ans[l-1] == -x { // if last and x is same with diff sign
				l--
			}
		}
	}

	return ans[:l]
}
