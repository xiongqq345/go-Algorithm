package coding

// 给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	return &TreeNode{
		Val:   nums[n/2],
		Left:  sortedArrayToBST(nums[:n/2]),
		Right: sortedArrayToBST(nums[n/2+1:]),
	}
}
