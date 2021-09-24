package jz_offerII

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

//多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
//
//给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/Qv1Da2
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	var cur, next *Node = root, nil
	for cur != nil {
		next = cur.Next
		if cur.Child != nil {
			child := cur.Child
			cur.Child, cur.Next, child.Prev = nil, cur.Child, cur
			for child.Next != nil {
				child = child.Next
			}
			child.Next = next
			if next != nil {
				next.Prev = child
			}
		}
		cur = cur.Next
	}
	return root
}
