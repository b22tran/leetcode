package q1_10

/** https://leetcode.com/problems/reverse-integer/description/
//main.go
	ans := q1_10.Reverse(-123)
	fmt.Printf("Answer: %v\n",ans)

// potential ways
	1. utilize multiplication/division(mod) to retrieve reversed values
	2. bitwise operation?? maybe not
	3. conv. to string, reverse, convert to int64 (> 32bit thres.) ? value : 0

	// threshold 2^32-1 = 2,147,483,647 for 32-bit signed integer
*/
func Reverse(x int) int {
	const maxInt = 1 <<31 - 1
	const minInt = -1 << 31
	reversed := 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	if reversed > maxInt || reversed < minInt {
		return 0
	}
	return reversed
}