package jz_offer

// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	var p, q int
	for p = i; postorder[p] < postorder[j]; p++ {
	}
	for q = p; postorder[q] > postorder[j]; q++ {
	}

	return q == j && recur(postorder, i, p-1) && recur(postorder, p, q-1)
}
