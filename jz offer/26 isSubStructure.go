package jz_offer

// 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
//
//B是A的子结构， 即 A中有出现和B相同的结构和节点值。
func isSubStructure(a *TreeNode, b *TreeNode) bool {
	if a == nil || b == nil {
		return false
	}
	if recur(a, b) {
		return true
	}
	return isSubStructure(a.Left, b) || isSubStructure(a.Right, b)
}

func recur(a, b *TreeNode) bool {
	if b == nil {
		return true
	}

	if a == nil || a.Val != b.Val {
		return false
	}
	return recur(a.Left, b.Left) && recur(a.Right, b.Right)
}
