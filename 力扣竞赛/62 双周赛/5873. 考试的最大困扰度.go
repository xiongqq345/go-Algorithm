package _2_双周赛

//一位老师正在出一场由 n 道判断题构成的考试，每道题的答案为 true （用 'T' 表示）或者 false （用 'F' 表示）。老师想增加学生对自己做出答案的不确定性，方法是 最大化 有 连续相同 结果的题数。（也就是连续出现 true 或者连续出现 false）。
//
//给你一个字符串 answerKey ，其中 answerKey[i] 是第 i 个问题的正确结果。除此以外，还给你一个整数 k ，表示你能进行以下操作的最多次数：
//
//每次操作中，将问题的正确答案改为 'T' 或者 'F' （也就是将 answerKey[i] 改为 'T' 或者 'F' ）。
//请你返回在不超过 k 次操作的情况下，最大 连续 'T' 或者 'F' 的数目。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximize-the-confusion-of-an-exam
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func maxConsecutiveAnswers(answerKey string, k int) int {
	nums1 := make([]int, len(answerKey))
	for i := range nums1 {
		if answerKey[i] == 'T' {
			nums1[i] = 1
		}
	}

	nums2 := make([]int, len(answerKey))
	for i := range nums2 {
		if answerKey[i] == 'F' {
			nums2[i] = 1
		}
	}

	return max(longestOnes(nums1, k), longestOnes(nums2, k))
}

func longestOnes(nums []int, k int) (ans int) {
	var left, lsum, rsum int
	for right, v := range nums {
		rsum += 1 - v
		for lsum < rsum-k {
			lsum += 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
