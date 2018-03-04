package q1_10

import (
	"math/big"
	"strconv"
	"strings"
	"fmt"
)
/**
// https://leetcode.com/problems/add-two-numbers/description/

Another approach is to add each node in tandem from l1, l2 & implement carry over
	Reason why this approach is slow
	- unnecessary usage of large integers thus resorting to math/big Int
	- using exp func on big numbers

	//main.go
	l1 := q1_10.CreateListNodeReversed([]string{"2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","9"})
	l2 := q1_10.CreateListNodeReversed([]string{"5","6","4","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","2","4","3","9","9","9","9"})

	//l1 := q1_10.CreateListNodeReversed([]string{"8","1"})
	//l2 := q1_10.CreateListNodeReversed([]string{"0"})

	ans := q1_10.AddTwoNumbers(l1,l2)
	fmt.Printf("Answer: %v",ans)
 */

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
	slice1 := JoinNumsFromListNode(l1, l1)
	slice2 := JoinNumsFromListNode(l2, l2)
	fmt.Printf("slice: %v %v\n", slice1, slice2)

	n1 := FormatToNumber(slice1) // converting the int array into int to retrieve sum
	n2 := FormatToNumber(slice2)
	n1.Add(&n1,&n2)
	fmt.Printf("numbers: %v %v\n", n1,n2)


	str := n1.String()
	strAry := strings.Split(str, "")
	return CreateListNodeReversed(strAry)
}

func JoinNumsFromListNode(l1 *ListNode, current *ListNode) (slice []int64){
	for l1 != nil { // implement channels here
		slice = append(slice, int64(current.Val))
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	return slice
}

func FormatToNumber(nums []int64) big.Int {
	var num, sig big.Int
	for i, v := range nums {
		val := big.NewInt(v)
		sig.Exp(big.NewInt(10 ),big.NewInt(int64(i)),nil)
		num.Add(&num, sig.Mul(val, &sig))
	}
	return num
}

func CreateListNodeReversed(strAry []string) *ListNode {
	length := len(strAry) - 1
	if length == -1 {
		return nil
	}
	num, _ := strconv.Atoi(strAry[length])
	fmt.Printf("create list node: %v %v\n", length, num)
	return &ListNode{
		Val:  num,
		Next: CreateListNodeReversed(strAry[0:length]),
	}
}
