package q1_10

import (
	"math"
	"strconv"
	"strings"
	"fmt"
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
	var slice1, slice2 []int
	var n1, n2 int

	current := l1
	slice1 = joinNumsFromListNode(l1, current)

	current = l2
	slice2 = joinNumsFromListNode(l2, current)

	n1 = formatToNumber(slice1) // converting the int array into int to retrieve sum
	n2 = formatToNumber(slice2)
	n1 += n2

	str := strconv.Itoa(n1)
	strAry := strings.Split(str, "")
	return createListNode(strAry)
}

func joinNumsFromListNode(l1 *ListNode, current *ListNode) (slice []int){
	for l1.Next != nil { // implement channels here
		slice = append(slice, current.Val)
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	return slice
}

func formatToNumber(nums []int) (num int) {
	for i, v := range nums {
		num += v * int(math.Pow(10, float64(len(nums)-1-i)))
	}
	return num
}

func createListNode(strAry []string) *ListNode {
	length := len(strAry) - 1
	if length == -1 {
		return nil
	}
	num, _ := strconv.Atoi(strAry[length])
	fmt.Printf("create list node: %v %v\n", length, num)
	return &ListNode{
		Val:  num,
		Next: createListNode(strAry[0:length]),
	}
}
