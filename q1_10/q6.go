package q1_10

import "fmt"

/** https://leetcode.com/problems/zigzag-conversion/description/
//main.go
	ans := q1_10.Convert("PAYPALISHIRING",3)
	fmt.Printf("Answer: %v\n",ans)

// potential ways
	1. O(n2) - map string into a 2d char array which will conform to the zz pattern
		then append char by reading from index by index
	2. O(n) - find a mathematical alg that can determine the position of each char based on pos of
		original string (weighted vals to char, then sort in order of weighted vals)
			- probs, if use weighted vals, then the weights itself may become tooooo weighted
*/

func Convert(s string, numRows int) string {
	convertedString := ""
	maxLen := len(s)
	zigzag := make([]string, numRows)
	counter := 0
	for counter<maxLen {
		for j:=0; j<numRows; j++{
			if counter>=maxLen{
				break
			}
			zigzag[j] += string(s[counter])
			fmt.Printf("s[counter]: %v zz: %v\n",string(s[counter]), zigzag)
			counter++
		}
		for k:=numRows-2; k>0; k--{
			if counter>=maxLen{
				break
			}
			fmt.Printf("s[counter]: %v \n",string(s[counter]))
			zigzag[k] += string(s[counter])
			counter++
		}
	}
	counter = 0
	for i:=0; i<numRows; i++{
		convertedString += zigzag[i]
	}
	return convertedString
}