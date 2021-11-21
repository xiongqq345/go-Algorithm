package tree

func postorderTraversal(root *TreeNode) (vals []int) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		vals = append(vals, node.Val)
	}
	postorder(root)
	return
}

func postorderTraversal(root *TreeNode) (vals []int) {

}
