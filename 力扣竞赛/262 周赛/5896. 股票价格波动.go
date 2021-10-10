package _62_周赛

import "github.com/emirpasic/gods/trees/redblacktree"

//给你一支股票价格的数据流。数据流中每一条记录包含一个 时间戳 和该时间点股票对应的 价格 。
//
//不巧的是，由于股票市场内在的波动性，股票价格记录可能不是按时间顺序到来的。某些情况下，有的记录可能是错的。如果两个有相同时间戳的记录出现在数据流中，前一条记录视为错误记录，后出现的记录 更正 前一条错误的记录。
//
//请你设计一个算法，实现：
//
//更新 股票在某一时间戳的股票价格，如果有之前同一时间戳的价格，这一操作将 更正 之前的错误价格。
//找到当前记录里 最新股票价格 。最新股票价格 定义为时间戳最晚的股票价格。
//找到当前记录里股票的 最高价格 。
//找到当前记录里股票的 最低价格 。
//请你实现 StockPrice 类：
//
//StockPrice() 初始化对象，当前无股票价格记录。
//void update(int timestamp, int price) 在时间点 timestamp 更新股票价格为 price 。
//int current() 返回股票 最新价格 。
//int maximum() 返回股票 最高价格 。
//int minimum() 返回股票 最低价格 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/stock-price-fluctuation
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type StockPrice struct {
	*redblacktree.Tree             // 把价格存到 multiset 中
	prices             map[int]int // 时间戳对应的价格
	now, cur           int
}

func Constructor() StockPrice {
	return StockPrice{redblacktree.NewWithIntComparator(), map[int]int{}, 0, 0}
}

func (s *StockPrice) Update(timestamp, price int) {
	if p := s.prices[timestamp]; p > 0 {
		s.remove(p) // 更新价格前先把旧的删掉
	}
	s.put(price)                // 记录价格
	s.prices[timestamp] = price // 记录时间戳对应价格
	if timestamp >= s.now {
		s.now, s.cur = timestamp, price // 更新最新时间及价格
	}
}

func (s *StockPrice) Current() int { return s.cur }
func (s *StockPrice) Maximum() int { return s.Right().Key.(int) }
func (s *StockPrice) Minimum() int { return s.Left().Key.(int) }

func (s *StockPrice) put(v int) {
	c := 0
	if cnt, has := s.Get(v); has {
		c = cnt.(int)
	}
	s.Put(v, c+1)
}

func (s *StockPrice) remove(v int) {
	if cnt, _ := s.Get(v); cnt.(int) > 1 {
		s.Put(v, cnt.(int)-1)
	} else {
		s.Remove(v)
	}
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
