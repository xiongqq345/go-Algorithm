package coding

//给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
	var ans []*ListNode
	q := []*TreeNode{tree}
	for i := 0; len(q) > 0; i++ {
		dummy := new(ListNode)
		p := dummy
		var nq []*TreeNode
		for _, node := range q {
			p.Next = &ListNode{Val: node.Val}
			p = p.Next
			if node.Left != nil {
				nq = append(nq, node.Left)
			}
			if node.Right != nil {
				nq = append(nq, node.Right)
			}
		}
		ans = append(ans, dummy.Next)
		q = nq
	}
	return ans
}
