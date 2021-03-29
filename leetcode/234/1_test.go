package _34

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	var head, p *ListNode

	for i := 0; i < 5; i++ {
		if head == nil {
			head = &ListNode{int(math.Abs(float64(i - 2))), nil}
			p = head
		} else {
			p.Next = &ListNode{int(math.Abs(float64(i - 2))), nil}
			p = p.Next
		}
	}
	p = head
	for p != nil {
		//fmt.Println(p.Val)
		p = p.Next
	}
	assert.Equal(t, true, isPalindrome(head))
}
