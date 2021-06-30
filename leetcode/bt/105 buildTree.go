package bt

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var r int
	for i, v := range inorder {
		if v == preorder[0] {
			r = i
			break
		}
	}

	root.Left = buildTree(preorder[1:r+1], inorder[:r])
	root.Right = buildTree(preorder[r+1:], inorder[r+1:])
	return root
}
