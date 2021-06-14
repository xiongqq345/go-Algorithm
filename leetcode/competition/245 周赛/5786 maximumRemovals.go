package _45_周赛

import (
	"sort"
)

func maximumRemovals(s string, p string, removable []int) int {
	return sort.Search(len(removable), func(h int) bool {
		mp := make(map[int]struct{})
		for i := 0; i <= h; i++ {
			mp[removable[i]] = struct{}{}
		}
		return !isSubsequence(p, s, mp)
	})
}

func isSubsequence(sub string, str string, delMp map[int]struct{}) bool {
	p, q := 0, 0
	for ; p < len(sub) && q < len(str); q++ {
		if _, ok := delMp[q]; ok {
			continue
		}
		if sub[p] == str[q] {
			p++
		}
	}
	return p == len(sub)
}
