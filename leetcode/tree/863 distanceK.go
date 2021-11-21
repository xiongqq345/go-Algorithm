package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。
//
//返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。
//
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	mp := make(map[int]*TreeNode)
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node.Left != nil {
			mp[node.Left.Val] = node
			helper(node.Left)
		}
		if node.Right != nil {
			mp[node.Right.Val] = node
			helper(node.Right)
		}
	}
	helper(root)

	var ans []int
	var findAns func(node, from *TreeNode, depth int)
	findAns = func(node, from *TreeNode, depth int) {
		if node == nil {
			return
		}
		depth++
		if depth == k {
			ans = append(ans, node.Val)
		}
		if node.Left != from {
			findAns(node.Left, node, depth)
		}
		if node.Right != from {
			findAns(node.Right, node, depth)
		}
		if mp[node.Val] != from {
			findAns(mp[node.Val], node, depth)
		}
	}
	findAns(target, target, -1)
	return ans
}
