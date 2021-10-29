package jz_offerII

import "github.com/emirpasic/gods/trees/redblacktree"

//请实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。
//
//MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end 时间内增加一个日程安排，注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。
//
//当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。
//
//每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。否则，返回 false 并且不要将该日程安排添加到日历中。
//
//请按照以下步骤调用 MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)

type MyCalendar struct {
	t *redblacktree.Tree
}

func Constructor() MyCalendar {
	t := redblacktree.NewWithIntComparator()
	t.Put(-1, -1) // 哨兵
	return MyCalendar{t}
}

func (c *MyCalendar) Book(start, end int) bool {
	floor, _ := c.t.Floor(start)
	if floor.Value.(int) > start { // [start,end) 左侧区间的右端点超过了 start
		return false
	}
	if it := c.t.IteratorAt(floor); it.Next() && it.Key().(int) < end { // [start,end) 右侧区间的左端点小于 end
		return false
	}
	c.t.Put(start, end) // 可以插入区间 [start,end)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
