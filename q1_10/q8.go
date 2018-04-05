package q1_10

import (
	"strings"
	"unicode"
	"fmt"
)

/** https://leetcode.com/problems/string-to-integer-atoi/description/
//main.go
	ans := q1_10.MyAtoi("2147483648")
	fmt.Printf("Answer: %v\n",ans)

// requirements
	1. trim pre-string whitespaces
	2. first non-w/s char can be +/- indicating pos/neg
	3. followed by as many numerical digits as possible
	4. ignore any non numerical data following
	5. if first seq non numeric, then no conv, return 0
	6. if out of range, return INT_MAX (2147483647) or INT_MIN (-2147483648)
*/

func MyAtoi(str string) int {
	const maxInt = 1 << 31 -1
	const minInt = -1 << 31
	str = strings.TrimLeft(str," ")
	signed := 1
	num := 0
	//fmt.Printf("bool: %v\n", !unicode.IsDigit(rune(str[0])))

	if str == ""{
		return 0
	}
	if str[0] == '+' {
		str = str[1:]
	}else if str[0] == '-'{
		signed = -1
		str = str[1:]
	}
	if str == ""{
		return 0
	}
	if !unicode.IsDigit(rune(str[0])){
		return 0
	}
	for i:=0; i<len(str); i++{
		fmt.Printf("rune: %v\n", str[i])
		if unicode.IsDigit(rune(str[i])){
			num = num * 10 + int(str[i]-'0')
		}else{
			break
		}
		if signed*num > maxInt {
			return maxInt
		}else if signed*num < minInt{
			return minInt
		}
	}
	return signed* num
}