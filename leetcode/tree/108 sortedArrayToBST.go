package tree

// 数组转二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return &TreeNode{
		nums[mid],
		helper(nums, l, mid-1),
		helper(nums, mid+1, r),
	}
}
