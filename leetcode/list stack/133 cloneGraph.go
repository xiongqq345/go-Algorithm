package list_stack

// 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
//
//图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if _, ok := visited[node]; ok {
			return visited[node]
		}
		n := &Node{
			Val:       node.Val,
			Neighbors: []*Node{},
		}
		visited[node] = n
		for _, neighbor := range node.Neighbors {
			n.Neighbors = append(n.Neighbors, dfs(neighbor))
		}
		return n
	}
	return dfs(node)
}

func dfs(node *Node) *Node {

}
