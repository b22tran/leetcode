package main

import (
	"fmt"
	"leetcode/q1_10"
)

func main() {
	l1 := &q1_10.ListNode{
		Val: 2,
		Next: &q1_10.ListNode{
			Val: 4,
			Next: &q1_10.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &q1_10.ListNode{
		Val: 5,
		Next: &q1_10.ListNode{
			Val: 6,
			Next: &q1_10.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	//fmt.Printf("Answer: %v %v",l1.Val, l2.Next.Next.Next)
	ans := q1_10.AddTwoNumbers(l1,l2)
	fmt.Printf("Answer: %v",ans)
}
