package tree

func inorderTraversal(root *TreeNode) (vals []int) {
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		vals = append(vals, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

func inorderTraversal2(root *TreeNode) (vals []int) {
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		vals = append(vals, node.Val)
		node = node.Right
	}
	return
}

//
//func inorderTraversal(root *TreeNode) (res []int) {
//	for root != nil {
//		if root.Left != nil {
//			// predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
//			predecessor := root.Left
//			for predecessor.Right != nil && predecessor.Right != root {
//				// 有右子树且没有设置过指向 root，则继续向右走
//				predecessor = predecessor.Right
//			}
//			if predecessor.Right == nil {
//				// 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
//				predecessor.Right = root
//				// 遍历左子树
//				root = root.Left
//			} else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
//				res = append(res, root.Val)
//				// 恢复原样
//				predecessor.Right = nil
//				// 遍历右子树
//				root = root.Right
//			}
//		} else { // 没有左子树
//			res = append(res, root.Val)
//			// 若有右子树，则遍历右子树
//			// 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
//			root = root.Right
//		}
//	}
//	return
//}
