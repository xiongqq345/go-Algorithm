package backtracking

import (
	"strconv"
	"strings"
)

//给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
//
//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
//例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。

func restoreIpAddresses(s string) []string {
	var ans []string
	var helper func([]string, int)
	helper = func(subRes []string, start int) {
		if len(subRes) == 4 && start == len(s) {
			ans = append(ans, strings.Join(subRes, "."))
			return
		}

		if len(subRes) == 4 && start < len(s) {
			return
		}
		for length := 1; length <= 3; length++ {
			if start+length-1 >= len(s) {
				return
			}
			if length != 1 && s[start] == '0' {
				return
			}
			str := s[start : start+length]
			if n, _ := strconv.Atoi(str); n > 255 {
				return
			}
			helper(append(subRes, str), start+length)
		}
	}
	helper([]string{}, 0)
	return ans
}
