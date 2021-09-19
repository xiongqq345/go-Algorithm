package _8_双周赛

// 给你一个下标从 0 开始的字符串 s ，你需要找到两个 不重叠的回文 子字符串，它们的长度都必须为 奇数 ，使得它们长度的乘积最大。
//
//更正式地，你想要选择四个整数 i ，j ，k ，l ，使得 0 <= i <= j < k <= l < s.length ，且子字符串 s[i...j] 和 s[k...l] 都是回文串且长度为奇数。s[i...j] 表示下标从 i 到 j 且 包含 两端下标的子字符串。
//
//请你返回两个不重叠回文子字符串长度的 最大 乘积。
//
//回文字符串 指的是一个从前往后读和从后往前读一模一样的字符串。子字符串 指的是一个字符串中一段连续字符。

func maxProduct(s string) int64 {
	ans := 1
	for i := 1; i < len(s)-1; i++ {
		ans = max(ans, longestPalindrome(s[:i])*longestPalindrome(s[i:]))
	}
	return int64(ans)
}

func longestPalindrome(s string) int {
	var ans int
	for i := range s {
		ans = max(ans, Palindrome(s, i, i))
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Palindrome(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
