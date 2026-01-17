package solutions

func climbStairs(n int) int {
	
	if n <= 3 {
		return n
	}    

	prev1 := 3
	prev2 := 2
	cur := 0

	for k := 3 ; k < n ; k++ {
		cur = prev1 + prev2
		prev2 = prev1
		prev1 = cur
	}

	return cur
}