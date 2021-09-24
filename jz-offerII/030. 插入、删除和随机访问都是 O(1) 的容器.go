package jz_offerII

import "math/rand"

//设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构：
//
//insert(val)：当元素 val 不存在时返回 true ，并向集合中插入该项，否则返回 false 。
//remove(val)：当元素 val 存在时返回 true ，并从集合中移除该项，否则返回 false 。
//getRandom：随机返回现有集合中的一项。每个元素应该有 相同的概率 被返回。

type RandomizedSet struct {
	set  map[int]int
	keys []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{set: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.set[val]; ok {
		return false
	}
	rs.set[val] = len(rs.keys)
	rs.keys = append(rs.keys, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (rs *RandomizedSet) Remove(val int) bool {
	if _, ok := rs.set[val]; !ok {
		return false
	}
	p := rs.set[val]
	rs.keys[p] = rs.keys[len(rs.keys)-1]
	rs.keys = rs.keys[:len(rs.keys)-1]
	rs.set[rs.keys[p]] = p
	delete(rs.set, val)
	return true
}

/** Get a random element from the set. */
func (rs *RandomizedSet) GetRandom() int {
	return rs.keys[rand.Intn(len(rs.keys))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
