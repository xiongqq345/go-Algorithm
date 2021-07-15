package coding

// 设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
//
//如果指定节点没有对应的“下一个”节点，则返回null。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var pre *TreeNode
	var inorder func(*TreeNode) *TreeNode
	inorder = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		l := inorder(node.Left)
		if l != nil {
			return l
		}
		if pre == p {
			return node
		} else {
			pre = node
		}
		return inorder(node.Right)
	}
	return inorder(root)
}
