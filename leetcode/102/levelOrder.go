package _02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	list := []*TreeNode{root}
	for i := 0; len(list) > 0; i++ {
		res = append(res, []int{})
		var newList []*TreeNode
		for j := 0; j < len(list); j++ {
			res[i] = append(res[i], list[j].Val)
			if list[j].Left != nil {
				newList = append(newList, list[j].Left)
			}
			if list[j].Right != nil {
				newList = append(newList, list[j].Right)
			}
		}
		list = newList
	}
	return res
}
