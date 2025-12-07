package solutions

// solution is not mine but I like it
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	j := 2
	for i := 2; i < len(nums); i++ {
		if nums[j-1] != nums[i] || nums[j-1] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
