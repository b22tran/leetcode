package q1_10


func LengthOfLongestSubstring(s string) int {
	longest := 0
	if len(s) > 0 {
		longest = 1
	}
	for i, initChar := range s {
		currChars := make(map[int32]int,0)
		currChars[initChar] = 0
		for _, ascii := range s[i+1:]{
			if _, ok := currChars[ascii]; ok{
				isLongest(currChars,&longest)
				break
			}
			currChars[ascii] = 0
			isLongest(currChars,&longest)
		}
	}
	return longest
}

func isLongest(currChars map[int32]int, longest *int){
	if len(currChars) > *longest{
		*longest = len(currChars)
	}
}