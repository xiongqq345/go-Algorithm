package jz_offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
