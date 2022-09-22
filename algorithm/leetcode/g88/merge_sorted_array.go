package g88

func merge(nums1 []int, m int, nums2 []int, n int) {
	for a, b, tail := m-1, n-1, m+n-1; a >= 0 || b >= 0; tail-- {
		var cur int
		if a == -1 {
			cur = nums2[b]
			b--
		} else if b == -1 {
			cur = nums1[a]
			a--
		} else if nums1[a] > nums2[b] {
			cur = nums1[a]
			a--
		} else {
			cur = nums2[b]
			b--
		}
		nums1[tail] = cur
	}
}
