package _65_周赛

import (
	"math"
)

//链表中的 临界点 定义为一个 局部极大值点 或 局部极小值点 。
//
//如果当前节点的值 严格大于 前一个节点和后一个节点，那么这个节点就是一个  局部极大值点 。
//
//如果当前节点的值 严格小于 前一个节点和后一个节点，那么这个节点就是一个  局部极小值点 。
//
//注意：节点只有在同时存在前一个节点和后一个节点的情况下，才能成为一个 局部极大值点 / 极小值点 。
//
//给你一个链表 head ，返回一个长度为 2 的数组 [minDistance, maxDistance] ，其中 minDistance 是任意两个不同临界点之间的最小距离，maxDistance 是任意两个不同临界点之间的最大距离。如果临界点少于两个，则返回 [-1，-1] 。
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	a, b, c := head, head.Next, head.Next.Next
	var first, last int
	minn := math.MaxInt64
	for i := 1; c != nil; a, b, c, i = b, c, c.Next, i+1 {
		if a.Val < b.Val && b.Val > c.Val || a.Val > b.Val && b.Val < c.Val {
			if last > 0 {
				minn = min(minn, i-last)
			}
			if first == 0 {
				first = i
			}
			last = i
		}
	}
	if first == last {
		return []int{-1, -1}
	}
	return []int{minn, last - first}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
