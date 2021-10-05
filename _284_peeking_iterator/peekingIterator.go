package _284_peeking_iterator

import "container/list"

// iterator definition
type Iterator struct {
	cur *list.Element
	end *list.Element
}

func (this *Iterator) hasNext() bool {
	return this.cur != this.end
}

func (this *Iterator) next() int {
	return this.next()
}

type PeekingIterator struct {
	iter     *Iterator
	_hasNext bool
	_next    int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, iter.hasNext(), iter.next()}
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
