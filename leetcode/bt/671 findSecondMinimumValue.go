package bt

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。
//
//更正式地说，root.val = min(root.left.val, root.right.val) 总成立。
//
//给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findSecondMinimumValue(root *TreeNode) int {
	ans := math.MaxInt32
	var f bool
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val > root.Val {
			if node.Val == math.MaxInt32 {
				f = true
			}
			ans = min(ans, node.Val)
		}
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)

	if ans == math.MaxInt32 && !f {
		return -1
	}
	return ans
}
