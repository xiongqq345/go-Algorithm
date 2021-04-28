package bt

import "sort"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int
	for node := root; node.Left != nil; node = node.Left {
		depth++
	}

	return sort.Search(1<<(depth+1), func(k int) bool {
		if k <= 1<<depth {
			return false
		}
		bits := 1 << (depth - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
