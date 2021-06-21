package 微爱思扣_以_Code_会友

import (
	"container/heap"
	"sort"
)

//有 N 位扣友参加了微软与力扣举办了「以扣会友」线下活动。主办方提供了 2*N 道题目，整型数组 questions 中每个数字对应了每道题目所涉及的知识点类型。
//若每位扣友选择不同的一题，请返回被选的 N 道题目至少包含多少种知识点类型。
//
//示例 1：
//
//输入：questions = [2,1,6,2]
//
//输出：1
//
//解释：有 2 位扣友在 4 道题目中选择 2 题。
//可选择完成知识点类型为 2 的题目时，此时仅一种知识点类型
//因此至少包含 1 种知识点类型。
//
//示例 2：
//
//输入：questions = [1,5,1,3,4,5,2,5,3,3,8,6]
//
//输出：2
//
//解释：有 6 位扣友在 12 道题目中选择题目，需要选择 6 题。
//选择完成知识点类型为 3、5 的题目，因此至少包含 2 种知识点类型。
type heapInt []int

//Less  小于就是小跟堆，大于号就是大根堆
func (h *heapInt) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *heapInt) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapInt) Len() int           { return len(*h) }
func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heapInt) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}
func (h *heapInt) Peek() int {
	return (*h)[0]
}

func halfQuestions(questions []int) int {
	n := len(questions) / 2
	mp := make(map[int]int)
	for _, v := range questions {
		mp[v]++
	}
	h := new(heapInt)
	for _, v := range mp {
		heap.Push(h, v)
	}
	var ans int
	for n > 0 {
		n -= mp[heap.Pop(h).(int)]
		ans++
	}
	return ans
}

func halfQuestions2(questions []int) int {
	n := len(questions) / 2
	mp := make(map[int]int)
	for _, v := range questions {
		mp[v]++
	}

	arr, i := make([]int, len(mp)), 0
	for _, v := range mp {
		arr[i] = v
		i++
	}
	sort.Ints(arr)
	var ans int
	for i = len(mp) - 1; i >= 0 && n > 0; i-- {
		n -= arr[i]
		ans++
	}
	return ans
}
