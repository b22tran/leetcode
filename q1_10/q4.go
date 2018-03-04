package q1_10

import "fmt"

/** https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 	//main.go
	nums1 := []int{1, 3}
	nums2 := []int{2}
	answer := findMedianSortedArrays(nums1,nums2)
	fmt.Printf("\n%v", answer)
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for _, v := range nums2 {
		nums1 = append(nums1, v)
	}
	nums1 = MergeSort(nums1)
	length := len(nums1) / 2
	if len(nums1)%2 == 0 {
		fmt.Printf("even: %v, %v, %v", nums1[length-1:length+1], float64(nums1[length-1]), float64(nums1[length]))
		return (float64(nums1[length-1]) + float64(nums1[length])) / 2.0
	}
	fmt.Printf("odd: %v", float64((nums1[length-1]+nums1[length+1])/2))
	return float64(nums1[length])
}

// merge sort from https://austingwalters.com/merge-sort-in-go-golang/
// Runs MergeSort algorithm on a slice single
func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
