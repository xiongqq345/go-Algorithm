package jz_offer

// 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal, idx := preorder[0], 0
	for i, v := range inorder {
		if v == rootVal {
			idx = i
		}
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:len(inorder[:idx])+1], inorder[:idx]),
		Right: buildTree(preorder[len(inorder[:idx])+1:], inorder[idx+1:]),
	}
}
