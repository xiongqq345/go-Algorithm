package jz_offerII

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
	node := &Node{Val: x}
	if aNode == nil {
		node.Next = node
		return node
	}
	if aNode.Next == aNode {
		aNode.Next = node
		node.Next = aNode
		return aNode
	}

	prev, cur := aNode, aNode.Next
	f := true
	for prev != aNode || f {
		f = false
		if prev.Val <= x && x <= cur.Val || (prev.Val > cur.Val && (x > prev.Val || x < cur.Val)) {
			prev.Next = node
			node.Next = cur
			return aNode
		}
		prev, cur = cur, cur.Next
	}
	aNode.Next, node.Next = node, aNode.Next
	return aNode
}
