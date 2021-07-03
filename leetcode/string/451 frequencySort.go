package string

import "bytes"

// 给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
func frequencySort(s string) string {
	var set [128]int
	var maxCnt int
	var ans []byte
	for _, v := range s {
		set[v]++
		maxCnt = max(maxCnt, set[v])
	}

	bucket := make([][]byte, maxCnt+1)
	for ch, v := range set {
		if v > 0 {
			bucket[v] = append(bucket[v], byte(ch))
		}
	}
	for i := maxCnt; i >= 0; i-- {
		for _, v := range bucket[i] {
			ans = append(ans, bytes.Repeat([]byte{v}, i)...)
		}
	}
	return string(ans)
}
