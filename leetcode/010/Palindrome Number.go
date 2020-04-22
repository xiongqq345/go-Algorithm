package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x != 0 && x%10 == 0 {
		return false
	}

	ori := x
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		rev = rev*10 + pop
	}
	return ori == rev
}
