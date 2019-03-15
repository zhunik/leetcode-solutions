package solutions

func SearchMatrix(matrix [][]int, target int) bool {

	for i := len(matrix) - 1; i >= 0; i-- {
		m := matrix[i]
		if len(m) == 0 {
			return false
		}
		if m[0] == target {
			return true
		} else if m[0] > target {
			continue
		} else if m[0] < target {
			for _, v := range m {
				if v == target {
					return true
				}
			}
		}
	}
	return false
}
