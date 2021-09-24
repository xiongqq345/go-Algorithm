package jz_offerII

//给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算滑动窗口里所有数字的平均值。
//
//实现 MovingAverage 类：
//
//MovingAverage(int size) 用窗口大小 size 初始化对象。
//double next(int val) 成员函数 next 每次调用的时候都会往滑动窗口增加一个整数，请计算并返回数据流中最后 size 个值的移动平均值，即滑动窗口里所有数字的平均值。

type MovingAverage struct {
	eles []int
	sum  int
	size int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (ma *MovingAverage) Next(val int) float64 {
	if len(ma.eles) == ma.size {
		ma.sum -= ma.eles[0]
		ma.eles = ma.eles[1:]
	}
	ma.sum += val
	ma.eles = append(ma.eles, val)
	return float64(ma.sum) / float64(len(ma.eles))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
