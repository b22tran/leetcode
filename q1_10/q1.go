package q1_10

import (
	"log"
)

/** https://leetcode.com/problems/two-sum/description/
//main.go
 nums := []int{2,5,5,11}
 twoSum := twoSum(nums, 10)
 fmt.Printf("%v", twoSum)
*/

// initially did a hybrid brute force, adapted their sol'n secondly
func twoSum(nums []int, target int) []int {
	nMap := make(map[int]int)
	for i, v := range nums {
		complement := target - nums[i]
		val, ok := nMap[complement]
		if ok {
			return []int{val, i}
		}
		nMap[v] = i
	}
	log.Fatal("No valid answer.")
	return nil
}
