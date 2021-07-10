package design

import "sort"

type TimeMap struct {
	mp map[string][]Value
}

type Value struct {
	value     string
	timestamp int
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		mp: make(map[string][]Value),
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.mp[key] = append(tm.mp[key], Value{
		value,
		timestamp,
	})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	list := tm.mp[key]
	i := sort.Search(len(list), func(i int) bool {
		return list[i].timestamp > timestamp
	})
	if i > 0 {
		return list[i-1].value
	}
	return ""
}
