package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/v3.0/luis/authoring"
)

//小美是美团仓库的管理员，她会根据单据的要求按顺序取出仓库中的货物，每取出一件货物后会把剩余货物重新堆放，使得自己方便查找。已知货物入库的时候是按顺序堆放在一起的。如果小美取出其中一件货物，则会把货物所在的一堆物品以取出的货物为界分成两堆，这样可以保证货物局部的顺序不变。
//已知货物最初是按 1~n 的顺序堆放的，每件货物的重量为 w[i] ,小美会根据单据依次不放回的取出货物。请问根据上述操作，小美每取出一件货物之后，重量和最大的一堆货物重量是多少？
//格式：
//
//
//输入：
//- 输入第一行包含一个正整数 n ，表示货物的数量。
//- 输入第二行包含 n 个正整数，表示 1~n 号货物的重量 w[i] 。
//- 输入第三行有 n 个数，表示小美按顺序取出的货物的编号，也就是一个 1~n 的全排列。
//输出：
//- 输出包含 n 行，每行一个整数，表示每取出一件货物以后，对于重量和最大的一堆货物，其重量和为多少。
//示例：
//
//
//输入：
//     5
//     3 2 4 4 5
//     4 3 5 2 1
//输出：
//     9
//     5
//     5
//     3
//     0
//
var sc, wr = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)

func scanString() string { sc.Scan(); return sc.Text() }
func scanRunes() []rune  { return []rune(scanString()) }
func scanInt() int       { a, _ := strconv.Atoi(scanString()); return a }
func scanInt64() int64   { a, _ := strconv.ParseInt(scanString(), 10, 64); return a }
func scanFloat() float64 { a, _ := strconv.ParseFloat(scanString(), 64); return a }

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func main() {
	n := scanInt()
	w := scanInts(n)
	take := scanInts(n)
	solution(w, take)
}

func solution(w, take []int) {
	s := make([]int, len(w)+1)
	var sum int
	for i := range w {
		sum += w[i]
		s[i+1] = sum
	}

}
