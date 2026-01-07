package solutions

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}
	if len(nums) == 0 {
        return res
    }
	if len(nums) == 1 {
		res = append(res, fmt.Sprintf("%d", nums[0]))
		return res
	}
	
	l := 0
	for r := 1; r < len(nums); r++ {
		if nums[r]-nums[r-1] != 1 {
			if l < r-1 {
				res = append(res, fmt.Sprintf("%d->%d", nums[l], nums[r-1]))
			} else {
				res = append(res, fmt.Sprintf("%d", nums[l]))
			}
			l = r
		} else if r == len(nums)-1 {
			res = append(res, fmt.Sprintf("%d->%d", nums[l], nums[r]))
		}

		if l == len(nums)-1 {
			res = append(res, fmt.Sprintf("%d", nums[l]))
		}
	}

	return res
}
