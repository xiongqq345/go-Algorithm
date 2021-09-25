package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//链接：https://ac.nowcoder.com/acm/contest/11179/A
//来源：牛客网
//
//牛牛很喜欢吃米饭，他的一顿饭是有 s 个米粒组成。一个偶然的机会，他找到了一条直线，这条直线上有 n 个格子，第一个格子有1个米粒，第二个格子有2个米粒，第三个格子有4个米粒……依次类推，第 n 个格子有 2n-1 个米粒。但由于有 k 个格子存在问题，会导致放在这些格子上的米粒全部丢失。现在想从剩下的 n-k 个格子中选出某些格子，拿走选中的格子上的全部米粒，问是否有恰好选出s个米粒的方案？
//输入描述:
//第一行三个整数 n 、 k 、s
//第二行 k 个整数，a_1a1,a_2a2,...,a_kak
//  表示有问题的格子的标号(格子标号从1开始)
//
//输出描述:
//能组成一个数量为 s 的米粒堆输出  YES，否则输出  NO
var sc, wr = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)

func scanString() string { sc.Scan(); return sc.Text() }
func scanRunes() []rune  { return []rune(scanString()) }
func scanInt() int       { a, _ := strconv.Atoi(scanString()); return a }
func scanInt64() int64   { a, _ := strconv.ParseInt(scanString(), 10, 64); return a }
func scanUint64() uint64 { a, _ := strconv.ParseUint(scanString(), 10, 64); return a }
func scanFloat() float64 { a, _ := strconv.ParseFloat(scanString(), 64); return a }
func pow2(a int) uint64  { return uint64(math.Pow(2, float64(a))) }

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}

func scanInts64(n int) []int64 {
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt64()
	}
	return a
}

func in(a []int, n int) bool {
	for _, v := range a {
		if v == n {
			return true
		}
	}
	return false
}

func main() {
	n := scanInt()
	k := scanInt()
	s := scanUint64()
	bads := scanInts(k)
	last := pow2(n - 1)
	for i := int(n) - 1; i >= 0; i-- {
		if in(bads, i+1) {
			last /= 2
			continue
		}
		if last <= s {
			s -= last
		}
		last /= 2
	}

	if s == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
