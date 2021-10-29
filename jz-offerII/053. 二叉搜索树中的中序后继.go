package jz_offerII

//给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。
//
//节点 p 的后继是值比 p.val 大的节点中键值最小的节点，即按中序遍历的顺序节点 p 的下一个节点。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/P5rCT8
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var f bool
	var ans *TreeNode
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		if f && ans == nil {
			ans = node
			return
		}
		if node == p {
			f = true
		}
		inorder(node.Right)
	}
	inorder(root)
	return ans
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var ans *TreeNode
	for root != nil {
		if root.Val > p.Val {
			ans = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return ans
}
