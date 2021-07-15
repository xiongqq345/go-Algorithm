package coding

// 动物收容所。有家动物收容所只收容狗与猫，且严格遵守“先进先出”的原则。在收养该收容所的动物时，收养人只能收养所有动物中“最老”（由其进入收容所的时间长短而定）的动物，或者可以挑选猫或狗（同时必须收养此类动物中“最老”的）。换言之，收养人不能自由挑选想收养的对象。请创建适用于这个系统的数据结构，实现各种操作方法，比如enqueue、dequeueAny、dequeueDog和dequeueCat。允许使用Java内置的LinkedList数据结构。
//
//enqueue方法有一个animal参数，animal[0]代表动物编号，animal[1]代表动物种类，其中 0 代表猫，1 代表狗。
//
//dequeue*方法返回一个列表[动物编号, 动物种类]，若没有可以收养的动物，则返回[-1,-1]。

type AnimalShelf struct {
	cats, dogs [][]int
}

func Constructor() AnimalShelf {
	return AnimalShelf{}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	if animal[1] == 0 {
		this.cats = append(this.cats, animal)
	} else {
		this.dogs = append(this.dogs, animal)
	}
}

func (this *AnimalShelf) DequeueAny() []int {
	if len(this.cats) == 0 {
		return this.DequeueDog()
	}
	if len(this.dogs) == 0 {
		return this.DequeueCat()
	}
	if this.cats[0][0] < this.dogs[0][0] {
		return this.DequeueCat()
	}
	return this.DequeueDog()
}

func (this *AnimalShelf) DequeueDog() []int {
	if len(this.dogs) == 0 {
		return []int{-1, -1}
	}
	as := this.dogs[0]
	this.dogs = this.dogs[1:]
	return as
}

func (this *AnimalShelf) DequeueCat() []int {
	if len(this.cats) == 0 {
		return []int{-1, -1}
	}
	as := this.cats[0]
	this.cats = this.cats[1:]
	return as
}
