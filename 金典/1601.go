package coding

// 编写一个函数，不用临时变量，直接交换numbers = [a, b]中a与b的值。
func swapNumbers(numbers []int) []int {
	numbers[0] ^= numbers[1]
	numbers[1] ^= numbers[0]
	numbers[0] ^= numbers[1]
	return numbers
}
