package bt

// 根据一棵树的中序遍历与后序遍历构造二叉树。
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	for i, v := range inorder {
		if v == postorder[len(postorder)-1] {
			return &TreeNode{
				Val:   postorder[len(postorder)-1],
				Left:  buildTree(inorder[:i], postorder[:i]),
				Right: buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
			}
		}
	}

	return nil
}
