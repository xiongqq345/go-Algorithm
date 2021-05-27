package linked_list

type MyLinkedList struct {
	head   *ListNode
	length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}
	p := this.head
	for index > 0 {
		p = p.Next
		index--
	}
	return p.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &ListNode{
		Val:  val,
		Next: this.head,
	}
	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.length == 0 {
		this.head = &ListNode{
			Val:  val,
			Next: nil,
		}
		this.length++
		return
	}
	p := this.head
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
	this.length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.length {
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	dummy := &ListNode{Next: this.head}
	for dummy != nil && index > 0 {
		dummy = dummy.Next
		index--
	}
	dummy.Next = &ListNode{
		Val:  val,
		Next: dummy.Next,
	}
	this.length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	}
	dummy := &ListNode{Next: this.head}
	p := dummy
	for p != nil && index > 0 {
		p = p.Next
		index--
	}
	p.Next = p.Next.Next
	this.head = dummy.Next
	this.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
