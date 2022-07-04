//
//
//

package list

type List_t[Value_t any] struct {
	Value Value_t
	prev  *List_t[Value_t]
	next  *List_t[Value_t]
	size  int
}

// for compatibility with go-circular
func New[Value_t any](...int) (self *List_t[Value_t]) {
	self = &List_t[Value_t]{}
	self.prev = self
	self.next = self
	return
}

func (it *List_t[Value_t]) cut_list() *List_t[Value_t] {
	it.prev.next = it.next
	it.next.prev = it.prev
	return it
}

func (it *List_t[Value_t]) set_before(at *List_t[Value_t]) *List_t[Value_t] {
	it.prev = at.prev
	at.prev.next = it
	at.prev = it
	it.next = at
	return it
}

func (it *List_t[Value_t]) set_after(at *List_t[Value_t]) *List_t[Value_t] {
	it.next = at.next
	at.next.prev = it
	at.next = it
	it.prev = at
	return it
}

func (self *List_t[Value_t]) PushFront(value Value_t) bool {
	(&List_t[Value_t]{Value: value}).set_after(self)
	self.size++
	return true
}

func (self *List_t[Value_t]) PushBack(value Value_t) bool {
	(&List_t[Value_t]{Value: value}).set_before(self)
	self.size++
	return true
}

func (self *List_t[Value_t]) PopFront() (v Value_t, ok bool) {
	if self.size == 0 {
		return
	}
	self.size--
	return self.next.cut_list().Value, true
}

func (self *List_t[Value_t]) PopBack() (v Value_t, ok bool) {
	if self.size == 0 {
		return
	}
	self.size--
	return self.prev.cut_list().Value, true
}

func (self *List_t[Value_t]) Front() *List_t[Value_t] {
	return self.next
}

func (self *List_t[Value_t]) Back() *List_t[Value_t] {
	return self.prev
}

func (self *List_t[Value_t]) End() *List_t[Value_t] {
	return self
}

func (self *List_t[Value_t]) Next() *List_t[Value_t] {
	return self.next
}

func (self *List_t[Value_t]) Prev() *List_t[Value_t] {
	return self.prev
}

func (self *List_t[Value_t]) Size() int {
	return self.size
}

func (self *List_t[Value_t]) RangeFront(f func(Value_t) bool) {
	for it := self.Front(); it != self.End(); it = it.Next() {
		if f(it.Value) == false {
			return
		}
	}
}

func (self *List_t[Value_t]) RangeBack(f func(Value_t) bool) {
	for it := self.Back(); it != self.End(); it = it.Prev() {
		if f(it.Value) == false {
			return
		}
	}
}

// linear if sorted before
func (self *List_t[Value_t]) InsertionSortFront(less Less[Value_t]) {
	for it1 := self.Front().Next(); it1 != self.End(); it1 = it1.Next() {
		for it2 := it1; it2.Prev() != self.End() && less(it2.Value, it2.Prev().Value); {
			it2.cut_list().set_before(it2.Prev())
		}
	}
}

// linear if sorted before
func (self *List_t[Value_t]) InsertionSortBack(less Less[Value_t]) {
	for it1 := self.Back().Prev(); it1 != self.End(); it1 = it1.Prev() {
		for it2 := it1; it2.Next() != self.End() && less(it2.Value, it2.Next().Value); {
			it2.cut_list().set_after(it2.Next())
		}
	}
}

type Less[Value_t any] func(a Value_t, b Value_t) bool
