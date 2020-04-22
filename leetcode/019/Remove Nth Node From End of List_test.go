package leetcode

import (
	"fmt"
	"testing"
)

func printNode(head *ListNode) {
	for head != nil {
		//fmt.Print(head.value, "\t")
		fmt.Println(head)
		head = head.Next
	}
	fmt.Println()
}

func Test_removeNthFromEnd(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 3
	node4 := new(ListNode)
	node4.Val = 4
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	printNode(removeNthFromEnd(node1, 1))
}
