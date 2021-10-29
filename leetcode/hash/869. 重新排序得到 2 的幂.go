package hash

//给定正整数 N ，我们按任何顺序（包括原始顺序）将数字重新排序，注意其前导数字不能为零。
//
//如果我们可以通过上述方式得到 2 的幂，返回 true；否则，返回 false。

func reorderedPowerOf2(n int) bool {
	powerOf2Digits := map[[10]int]bool{}
	for i := 1; i <= 1e9; i <<= 1 {
		powerOf2Digits[countDigits(i)] = true
	}

	return powerOf2Digits[countDigits(n)]
}

func countDigits(n int) (cnt [10]int) {
	for n > 0 {
		cnt[n%10]++
		n /= 10
	}
	return
}
