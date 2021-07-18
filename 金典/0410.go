package coding

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 检查子树。你有两棵非常大的二叉树：T1，有几万个节点；T2，有几万个节点。设计一个算法，判断 T2 是否为 T1 的子树。
//
//如果 T1 有这么一个节点 n，其子树与 T2 一模一样，则 T2 为 T1 的子树，也就是说，从节点 n 处把树砍断，得到的树与 T2 完全相同。
//
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil || t2 == nil {
		return false
	}
	if check(t1, t2) {
		return true
	}
	return checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}

func check(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil || a.Val != b.Val {
		return false
	}
	return check(a.Left, b.Left) && check(a.Right, b.Right)
}
