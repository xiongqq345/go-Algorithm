package main

import "fmt"

//链接：https://ac.nowcoder.com/acm/contest/11179/B
//来源：牛客网
//
//牛牛吃饱之后觉的非常口渴，于是他找出了他最喜欢的cocacola！
//牛牛的强迫症很强，虽然都是cocacola，但如果其中的某些字母的顺序颠倒或位置互换，他就不想去喝它。
//为了尽快喝到cocacola，他把这个问题交给了你，希望你能告诉他最少需要交换多少次字符位置(每次仅可交换一个字母)可以得到cocacola
//输入描述:
//一行长度为8的字符串且保证有解
//输出描述:
//一个数字，表示得到 cocacola 的最少交换次数

func main() {
	ori := [8]byte{'c', 'o', 'c', 'a', 'c', 'o', 'l', 'a'}
	var s string
	fmt.Scan(&s)
	str := []byte(s)

	for i := 0; i < 8; i++ {
		if str[i] != ori[i] {

		}
	}
}
