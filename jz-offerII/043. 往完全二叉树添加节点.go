package jz_offerII

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//完全二叉树是每一层（除最后一层外）都是完全填充（即，节点数达到最大，第 n 层有 2n-1 个节点）的，并且所有的节点都尽可能地集中在左侧。
//
//设计一个用完全二叉树初始化的数据结构 CBTInserter，它支持以下几种操作：
//
//CBTInserter(TreeNode root) 使用根节点为 root 的给定树初始化该数据结构；
//CBTInserter.insert(int v)  向树中插入一个新节点，节点类型为 TreeNode，值为 v 。使树保持完全二叉树的状态，并返回插入的新节点的父节点的值；
//CBTInserter.get_root() 将返回树的根节点。
type CBTInserter struct {
	p []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	cbt := &CBTInserter{[]*TreeNode{root}}
	for i := 0; i < len(cbt.p); i++ {
		if cbt.p[i].Left != nil {
			cbt.p = append(cbt.p, cbt.p[i].Left)
		}
		if cbt.p[i].Right != nil {
			cbt.p = append(cbt.p, cbt.p[i].Right)
		}
	}
	return *cbt
}

func (cbt *CBTInserter) Insert(v int) int {
	n := len(cbt.p)
	father := cbt.p[(n-1)>>1]
	cbt.p = append(cbt.p, &TreeNode{Val: v})
	if n&1 == 1 {
		father.Left = cbt.p[n]
	} else {
		father.Right = cbt.p[n]
	}
	return father.Val
}

func (cbt *CBTInserter) Get_root() *TreeNode {
	return cbt.p[0]
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
