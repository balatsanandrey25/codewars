package main

func main() {
	nums1 := []int{1, 2, 2, 0, 0, 0}
	m := 3
	nums2 := []int{2, 2, 6}
	n := 3
	merge(nums1, m, nums2, n)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	f1 := m - 1
	f2 := n - 1
	full := m + n - 1
	for f1 >= 0 && f2 >= 0 {
		if nums1[f1] >= nums2[f2] {
			nums1[full] = nums1[f1]
			f1--
		} else {
			nums1[full] = nums2[f2]
			f2--
		}
		full--
	}
	for f2 >= 0 {
		nums1[full] = nums2[f2]
		f2--
		full--
	}

}
