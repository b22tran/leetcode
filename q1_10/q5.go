package q1_10

/** https://leetcode.com/problems/longest-palindromic-substring/description/
//main.go
	ans := q1_10.LongestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth")
	fmt.Printf("Answer: %v\n",ans)
*/

func LongestPalindrome(s string) string {
	// instantiate return vars
	longestLength := 0
	longestPalindrome := ""
	stringLength := len(s)
	for i:=0; i < len(s); i++{
		// instantiate temp vars
		currPalindromeLen := 0
		currPalindrome := ""
		j := stringLength
		for j>i{
			if len(s[i:j]) < longestLength{
				break // doesn't check string that are shorter than current longest
			}
			//fmt.Printf("s[i:j]: %v j: %v \n",s[i:j],j)
			currPalindromeLen, currPalindrome= checkPalindrome(s[i:j])
			j-- // reduce the window
			if currPalindromeLen > longestLength{
				longestLength = currPalindromeLen
				longestPalindrome = currPalindrome
			}
		}
	}
	// fmt.Printf(longestPalindrome)
	return longestPalindrome
}

func checkPalindrome(substring string) (int, string){
	strLength := len(substring)
		for i := 0; i < (strLength/2); i++{
			//fmt.Printf(" ss1: %v ss2: %v\n",substring[i],substring[strLength-i-1])
			if substring[i] == substring[strLength-i-1]{
				continue
			}else{
				return 0, ""
			}
		}
	return strLength, substring
}