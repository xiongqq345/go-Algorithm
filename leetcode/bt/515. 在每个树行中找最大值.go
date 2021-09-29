package bt

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	var ans []int
	for len(q) > 0 {
		var nq []*TreeNode
		max := q[0].Val
		for _, node := range q {
			if node == nil {
				continue
			}
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				nq = append(nq, node.Left)
			}
			if node.Right != nil {
				nq = append(nq, node.Right)
			}
		}
		q = nq
		ans = append(ans, max)
	}
	return ans
}
