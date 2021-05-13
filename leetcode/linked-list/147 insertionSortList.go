package linked_list

//插入排序算法：
//
//插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
//每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
//重复直到所有输入数据插入完为止。
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	lastSorted, cur := head, head.Next
	for cur != nil {
		if cur.Val >= lastSorted.Val {
			lastSorted = cur
		} else {
			pre := dummy
			for pre.Next.Val < cur.Val {
				pre = pre.Next
			}
			lastSorted.Next = cur.Next
			pre.Next, cur.Next = cur, pre.Next
		}
		cur = lastSorted.Next
	}
	return dummy.Next
}
