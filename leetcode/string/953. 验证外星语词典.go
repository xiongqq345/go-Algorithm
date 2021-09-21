package string

import "sort"

//某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。
//
//给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；否则，返回 false。
//
func isAlienSorted(words []string, order string) bool {
	mp := make(map[byte]int)
	for i, ch := range order {
		mp[byte(ch)] = i
	}
	tmp := make([]string, len(words))
	copy(tmp, words)
	sort.Slice(tmp, func(i, j int) bool {
		for idx := 0; idx < len(tmp[i]) && idx < len(tmp[j]); idx++ {
			if mp[tmp[i][idx]] != mp[tmp[j][idx]] {
				return mp[tmp[i][idx]] < mp[tmp[j][idx]]
			}
		}
		return len(tmp[i]) < len(tmp[j])
	})
	for i := range tmp {
		if tmp[i] != words[i] {
			return false
		}
	}
	return true
}
