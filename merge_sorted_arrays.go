package solutions


func mergeSortedArrays(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 {
		return
	}
	current := m+n-1
	l := max(0, m-1)
	r := n-1
    
    for r >= 0 {
     	if l < 0 || nums2[r] > nums1[l] {
     		nums1[current] = nums2[r]
     		r -= 1
     	} else {
     		nums1[current] = nums1[l]
     		nums1[l] = nums2[r]
     		l -= 1
     	}
     	
     	current -= 1     	
    }
}