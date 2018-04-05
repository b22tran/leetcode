package q1_10

import "fmt"

/** https://leetcode.com/problems/palindrome-number/description/
//main.go
	ans := q1_10.IsPalindrome(-12132121)
	fmt.Printf("Answer: %v\n",ans)

// possible approaches
	1. mod, place int in slice, compare
	1a. get len of num, mod, place int in slice, when halfway, compare for pali
	2. conv to str, then use palindrome check from a modified q5
*/

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	ints := make([]int, 0)
	for x != 0 {
		ints = append(ints, x%10)
		x = x / 10
	}
	length := len(ints)
	for i := 0; i < length/2; i++ {
		fmt.Printf("v1: %v v2: %v\n", ints[i], ints[length-i-1])
		if ints[i] != ints[length-i-1] {
			return false
		}
	}
	return true
}
