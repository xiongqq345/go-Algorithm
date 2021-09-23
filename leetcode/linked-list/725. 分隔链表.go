package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//给定一个头结点为 root 的链表, 编写一个函数以将链表分隔为 k 个连续的部分。
//
//每部分的长度应该尽可能的相等: 任意两部分的长度差距不能超过 1，也就是说可能有些部分为 null。
//
//这k个部分应该按照在链表中出现的顺序进行输出，并且排在前面的部分的长度应该大于或等于后面的长度。
//
//返回一个符合上述规则的链表的列表。
//
//举例： 1->2->3->4, k = 5 // 5 结果 [ [1], [2], [3], [4], null ]
func splitListToParts(head *ListNode, k int) []*ListNode {
	var l int
	for p := head; p != nil; p = p.Next {
		l++
	}
	quo, rem := l/k, l%k
	var ans []*ListNode
	var cnt int
	var tmp *ListNode
	var rest int
	if rem > 0 {
		rest = 1
	}
	for p := head; len(ans) < k; {
		if tmp == nil {
			tmp = p
		}
		var next *ListNode
		if p != nil {
			next = p.Next
		}
		cnt++
		if cnt >= quo+rest {
			if p != nil {
				p.Next = nil
			}
			ans = append(ans, tmp)
			cnt = 0
			tmp = nil
			if rem > 1 {
				rem--
			} else {
				rest = 0
			}
		}
		p = next
	}
	return ans
}
