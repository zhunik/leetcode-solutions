package solutions

func TwoSum(nums []int, target int) []int {
	d := make(map[int]int)

	for i, n := range nums {

		key := target - n
		elem, ok := d[key]

		if ok {
			return []int{elem, i}
		}

		d[n] = i

	}

	return []int{0, 0}
}
