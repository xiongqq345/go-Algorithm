package linked_list

//请编写一个函数，用于 删除单链表中某个特定节点 。在设计函数时需要注意，你无法访问链表的头节点 head ，只能直接访问 要被删除的节点 。
//
//题目数据保证需要删除的节点 不是末尾节点 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/delete-node-in-a-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
