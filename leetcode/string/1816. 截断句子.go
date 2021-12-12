package string

import (
	"strings"
)

// 句子 是一个单词列表，列表中的单词之间用单个空格隔开，且不存在前导或尾随空格。每个单词仅由大小写英文字母组成（不含标点符号）。
//
// 例如，"Hello World"、"HELLO" 和 "hello world hello world" 都是句子。
// 给你一个句子 s​​​​​​ 和一个整数 k​​​​​​ ，请你将 s​​ 截断 ​，​​​使截断后的句子仅含 前 k​​​​​​ 个单词。返回 截断 s​​​​​​ 后得到的句子。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/truncate-sentence
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func truncateSentence(s string, k int) string {
	return strings.Join(strings.Fields(s)[:k], " ")
}
