package list_stack

type MovingAverage struct {
	queue         []int
	size, element int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		queue: make([]int, size),
		size:  size,
	}
}

func (mv *MovingAverage) Next(val int) float64 {
	if mv.element < mv.size {
		mv.queue[mv.element] = val
		mv.element++
	} else {
		mv.queue = mv.queue[1:]
		mv.queue = append(mv.queue, val)
	}

	var sum int
	for i := 0; i < mv.element; i++ {
		sum += mv.queue[i]
	}
	return float64(sum) / float64(mv.element)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
