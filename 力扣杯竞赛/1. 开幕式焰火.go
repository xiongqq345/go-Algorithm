package 力扣杯竞赛

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numColor(root *TreeNode) int {
	set := make(map[int]struct{})
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		set[node.Val] = struct{}{}
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return len(set)
}
