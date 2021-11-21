package tree

import "math"

// 最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	t := math.MaxInt32
	if root.Left != nil {
		t = min(minDepth(root.Left), t)
	}
	if root.Right != nil {
		t = min(minDepth(root.Right), t)
	}
	return t + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDepth2(root *TreeNode) int {
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		var nq []*TreeNode
		for j := 0; j < len(q); j++ {
			if q[j] == nil {
				continue
			}
			if q[j].Left == nil && q[j].Right == nil {
				return i + 1
			}
			nq = append(nq, q[j].Left, q[j].Right)
		}
		q = nq
	}
	return 0
}
