package q261_270

/** https://leetcode.com/problems/missing-number/description/
// possible approaches
	1. merge sort than find missing num
	2. hashmap of (all nums, bool), the key with false value is missing
	3. add all nums then subtract the difference
*/

func MissingNumber(nums []int) int {
	totalSum, inputSum := len(nums), 0
	for i := 0; i< len(nums); i++ {
		totalSum, inputSum = totalSum + i, inputSum + nums[i]
	}
	return totalSum - inputSum
}