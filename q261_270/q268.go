package q261_270

/** https://leetcode.com/problems/missing-number/description/
//main.go
	ans := q1_10.IsPalindrome(-12132121)
	fmt.Printf("Answer: %v\n",ans)

// possible approaches
	1.
*/

func MissingNumber(nums []int) int {
	totalSum, inputSum := len(nums), 0
	for i := 0; i< len(nums); i++ {
		totalSum, inputSum = totalSum + i, inputSum + nums[i]
	}
	return totalSum - inputSum
}