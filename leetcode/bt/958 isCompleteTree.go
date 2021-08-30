package bt

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 给定一个二叉树，确定它是否是一个完全二叉树。
//
//百度百科中对完全二叉树的定义如下：
//
//若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~ 2h 个节点。）
//
func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	var end bool
	for len(q) > 0 {
		var nq []*TreeNode
		for _, node := range q {
			if node == nil {
				end = true
			} else {
				if end {
					return false
				}
				nq = append(nq, node.Left)
				nq = append(nq, node.Right)
			}
		}
		q = nq
	}
	return true
}
