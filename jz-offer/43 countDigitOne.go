package jz_offer

// 输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
//
//例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

// todo timeout
func countDigitOne(n int) int {
	var count int
	for i := 1; i <= n; i++ {
		num := i
		for num > 0 {
			if num%10 == 1 {
				count++
			}
			num /= 10
		}
	}
	return count
}
