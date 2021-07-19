package bt

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
func findTarget(root *TreeNode, k int) bool {
	var arr []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		arr = append(arr, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	i, j := 0, len(arr)-1
	for i < j {
		if arr[i]+arr[j] < k {
			i++
		} else if arr[i]+arr[j] > k {
			j--
		} else {
			return true
		}
	}
	return false
}
