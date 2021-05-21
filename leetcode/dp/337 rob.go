package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
//计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

// 后序遍历 y,n代表偷和不偷
func rob3(root *TreeNode) int {
	var postorder func(*TreeNode) (y, n int)
	postorder = func(node *TreeNode) (y, n int) {
		if node == nil {
			return
		}
		ly, ln := postorder(node.Left)
		ry, rn := postorder(node.Right)
		return node.Val + ln + rn, max(ly, ln) + max(ry, rn)
	}
	return max(postorder(root))
}

// 暴力递归
//func rob(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	money := root.Val
//	if root.Left != nil {
//		money += rob(root.Left.Left) + rob(root.Left.Right)
//	}
//	if root.Right != nil {
//		money += rob(root.Right.Left) + rob(root.Right.Right)
//	}
//
//	return max(money, rob(root.Left)+rob(root.Right))
//}
