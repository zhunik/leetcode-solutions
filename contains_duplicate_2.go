package solutions


func containsNearbyDuplicate(nums []int, k int) bool {
	h := map[int]int{}

	abs := func(x int) int {
		if x < 0 {
			return -x 
		}
		return x
	}

	for i,v := range nums {
		j, ok := h[v]
		if ok && abs(j - i) <= k {
			return true
		}
		
		h[v] = i
	}

	return false
}