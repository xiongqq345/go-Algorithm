package jz_offerII

//写一个 RecentCounter 类来计算特定时间范围内最近的请求。
//
//请实现 RecentCounter 类：
//
//RecentCounter() 初始化计数器，请求数为 0 。
//int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。
//保证 每次对 ping 的调用都使用比之前更大的 t 值。
//

type RecentCounter struct {
	reqs []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.reqs = append(rc.reqs, t)
	for i, v := range rc.reqs {
		if v >= t-3000 {
			rc.reqs = rc.reqs[i:]
			break
		}
	}
	return len(rc.reqs)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
