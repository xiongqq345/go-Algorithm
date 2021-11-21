package tree

//在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
//
//如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
//
//我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
//
//只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。

var (
	depth  map[int]int
	parent map[int]int
)

func isCousins(root *TreeNode, x int, y int) bool {
	depth, parent = make(map[int]int), make(map[int]int)
	dfs(root, root.Val, 1)
	return depth[x] == depth[y] && parent[x] != parent[y]
}

func dfs(node *TreeNode, par, deep int) {
	if node == nil {
		return
	}
	depth[node.Val] = deep
	parent[node.Val] = par
	dfs(node.Left, node.Val, deep+1)
	dfs(node.Right, node.Val, deep+1)
}
