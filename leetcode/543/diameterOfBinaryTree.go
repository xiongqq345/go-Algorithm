package _43

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	depth(root, &ans)
	return ans
}

func depth(n *TreeNode, ans *int) int {
	if n == nil {
		return 0
	}
	l := depth(n.Left, ans)
	r := depth(n.Right, ans)
	if l+r > *ans {
		*ans = l + r
	}
	return max(l, r) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
