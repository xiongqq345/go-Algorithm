package _48

import (
	"fmt"
	"testing"
)

func Test_sortList(t *testing.T) {
	n1 := &ListNode{Val: 4}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 3}
	n1.Next, n2.Next, n3.Next = n2, n3, n4
	n := sortList(n1)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
