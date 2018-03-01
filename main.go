package main

import (
	"fmt"
	"leetcode/q1_10"
)

func main() {
	l1 := q1_10.CreateListNodeReversed([]string{"3","4","2"})
	l2 := q1_10.CreateListNodeReversed([]string{"4","6","5"})


	//l1 := q1_10.CreateListNodeReversed([]string{"8","1"})
	//l2 := q1_10.CreateListNodeReversed([]string{"0"})


	ans := q1_10.AddTwoNumbers(l1,l2)
	fmt.Printf("Answer: %v",ans)
}
