package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则，返回 false 。
//
//二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/subtree-of-another-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}
	if subTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func subTree(t1, t2 *TreeNode) bool {
	if t2 == nil && t1 == nil {
		return true
	}
	if t2 == nil || t1 == nil || t1.Val != t2.Val {
		return false
	}
	return subTree(t1.Left, t2.Left) && subTree(t1.Right, t2.Right)
}
