package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

//One day, Ahmed_Hossam went to Hemose and said "Let's solve a gym contest!". Hemose didn't want to do that, as he was playing Valorant, so he came up with a problem and told it to Ahmed to distract him. Sadly, Ahmed can't solve it... Could you help him?
//
//There is an Agent in Valorant, and he has 𝑛 weapons. The 𝑖-th weapon has a damage value 𝑎𝑖, and the Agent will face an enemy whose health value is 𝐻.
//
//The Agent will perform one or more moves until the enemy dies.
//
//In one move, he will choose a weapon and decrease the enemy's health by its damage value. The enemy will die when his health will become less than or equal to 0. However, not everything is so easy: the Agent can't choose the same weapon for 2 times in a row.
//
//What is the minimum number of times that the Agent will need to use the weapons to kill the enemy?
//
//Input
//Each test contains multiple test cases. The first line contains the number of test cases 𝑡 (1≤𝑡≤105). Description of the test cases follows.
//
//The first line of each test case contains two integers 𝑛 and 𝐻 (2≤𝑛≤103,1≤𝐻≤109) — the number of available weapons and the initial health value of the enemy.
//
//The second line of each test case contains 𝑛 integers 𝑎1,𝑎2,…,𝑎𝑛 (1≤𝑎𝑖≤109) — the damage values of the weapons.
//
//It's guaranteed that the sum of 𝑛 over all test cases doesn't exceed 2⋅105.
//
//Output
//For each test case, print a single integer — the minimum number of times that the Agent will have to use the weapons to kill the enemy.
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
	sc.Split(bufio.ScanWords)
	for t := scanInt(); t > 0; t-- {
		n, h := scanInt(), scanInt()
		a := scanInts(n)
		sort.Ints(a)
		times := h / (a[n-1] + a[n-2]) * 2
		if h%(a[n-1]+a[n-2]) > 0 {
			if h%(a[n-1]+a[n-2])%a[n-1] == 0 {
				times++
			} else {
				times += h%(a[n-1]+a[n-2])/a[n-1] + 1
			}
		}
		fmt.Println(times)
	}
}
