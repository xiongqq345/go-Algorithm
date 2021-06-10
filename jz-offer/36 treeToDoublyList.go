package jz_offer

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var head, pre *TreeNode
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if pre == nil {
			head = node
		} else {
			node.Left = pre
			pre.Right = node
		}
		pre = node
		inorder(node.Right)
	}
	inorder(root)
	head.Left = pre
	pre.Right = head
	return head
}
