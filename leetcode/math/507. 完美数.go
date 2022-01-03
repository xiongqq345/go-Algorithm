package math

//对于一个 正整数，如果它和除了它自身以外的所有 正因子 之和相等，我们称它为 「完美数」。
//
//给定一个 整数 n， 如果是完美数，返回 true，否则返回 false
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	for d := 2; d*d <= num; d++ {
		if num%d == 0 {
			sum += d
			if d*d < num {
				sum += num / d
			}
		}
	}
	return sum == num
}
