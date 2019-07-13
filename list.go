//
//
//

package list

type List_t struct {
	Value interface{}
	prev * List_t
	next * List_t
	size int
}

type Less_t interface {
	Less(a interface{}, b interface{}) bool
}

func New() (self * List_t) {
	self = &List_t{}
	self.prev = self
	self.next = self
	return
}

func cut_list(it * List_t) * List_t {
	it.prev.next = it.next
	it.next.prev = it.prev
	return it
}

func set_before(it * List_t, at * List_t) * List_t {
	it.prev = at.prev
	at.prev.next = it
	at.prev = it
	it.next = at
	return it
}

func set_after(it * List_t, at * List_t) * List_t {
	it.next = at.next
	at.next.prev = it
	at.next = it;
	it.prev = at
	return it
}

func (self * List_t) PushFront(value interface{}) {
	set_after(&List_t{Value: value}, self)
	self.size++
}

func (self * List_t) PushBack(value interface{}) {
	set_before(&List_t{Value: value}, self)
	self.size++
}

func (self * List_t) PopFront() (interface {}) {
	self.size--
	return cut_list(self.next).Value
}

func (self * List_t) PopBack() (interface {}) {
	self.size--
	return cut_list(self.prev).Value
}

func (self * List_t) Front() * List_t {
	return self.next
}

func (self * List_t) Back() * List_t {
	return self.prev
}

func (self * List_t) End() * List_t {
	return self
}

func (self * List_t) Next() * List_t {
	return self.next
}

func (self * List_t) Prev() * List_t {
	return self.prev
}

func (self * List_t) Size() int {
	return self.size
}

// takes linear time if sorted before
func (self * List_t) InsertionSortFront(less Less_t) {
	for it1 := self.Front().Next(); it1 != self.End(); it1 = it1.Next() {
		for it2 := it1; it2.Prev() != self.End() && less.Less(it2.Value, it2.Prev().Value); {
			set_before(cut_list(it2), it2.Prev())
		}
	}
}

// takes linear time if sorted before
func (self * List_t) InsertionSortBack(less Less_t) {
	for it1 := self.Back().Prev(); it1 != self.End(); it1 = it1.Prev() {
		for it2 := it1; it2.Next() != self.End() && less.Less(it2.Value, it2.Next().Value); {
			set_after(cut_list(it2), it2.Next())
		}
	}
}

type reverse_t struct {
	Less_t
}

func (self * reverse_t) Less(a interface{}, b interface{}) bool {
	return self.Less_t.Less(b, a)
}

func Reverse(less Less_t) Less_t {
	return &reverse_t{less}
}
