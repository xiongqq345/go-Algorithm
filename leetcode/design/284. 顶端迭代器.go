package design

//请你设计一个迭代器，除了支持 hasNext 和 next 操作外，还支持 peek 操作。
//
//实现 PeekingIterator 类：
//
//PeekingIterator(int[] nums) 使用指定整数数组 nums 初始化迭代器。
//int next() 返回数组中的下一个元素，并将指针移动到下个元素处。
//bool hasNext() 如果数组中存在下一个元素，返回 true ；否则，返回 false 。
//int peek() 返回数组中的下一个元素，但 不 移动指针。
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */
type Iterator interface {
	next() int
	hasNext() bool
}

type PeekingIterator struct {
	iter     Iterator
	_hasNext bool
	_next    int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{*iter, (*iter).hasNext(), (*iter).next()}
}

func (it *PeekingIterator) hasNext() bool {
	return it._hasNext
}

func (it *PeekingIterator) next() int {
	ret := it._next
	it._hasNext = it.iter.hasNext()
	if it._hasNext {
		it._next = it.iter.next()
	}
	return ret
}

func (it *PeekingIterator) peek() int {
	return it._next
}
