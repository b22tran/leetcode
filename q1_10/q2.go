package q1_10

import (
	"math"
	"strconv"
	"strings"
)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	slice1 := make([]int,0)
	slice2 := make([]int,0)
	var n1,n2 int
	var l3 *ListNode
	current := l1
	for l1.Next != nil{ // implement channels here
		slice1 = append(slice1, current.Val)
		current = current.Next
	}
	current = l2
	for l2.Next != nil{
		slice2 = append(slice2, current.Val)
		current = current.Next
	}
	for i,v := range slice1{
		n1 += v * int(math.Pow(10,float64(len(slice1)-i)))
	}
	for i,v := range slice2{
		n2 += v * int(math.Pow(10,float64(len(slice2)-i)))
	}
	n1 +=n2
	str := strconv.Itoa(n1)
	strAry := strings.Split(str,"")
	current = l3
	for _,v := range strAry{
		current.Val,_ = strconv.Atoi(v)
		var lnode *ListNode
		current.Next = lnode

	}
	return l3
}
