package list

import (
	"fmt"
	"testing"
)

func Example_list1() {
	l := New[string]()
	l.PushBack("ddddd")
	l.PushFront("ccccc")
	l.PushBack("aaaaa")
	l.PushFront("bbbbb")

	l.RangeFront(func(v string) bool {
		fmt.Printf("%v\n", v)
		return true
	})
	fmt.Printf("%v\n", l.Size())
	// Output:
	// bbbbb
	// ccccc
	// ddddd
	// aaaaa
	// 4
}

func LessStr(a string, b string) bool {
	return a < b
}

func Example_list2() {
	l := New[string]()
	l.PushBack("ddddd")
	l.PushFront("ccccc")
	l.PushBack("aaaaa")
	l.PushFront("bbbbb")

	l.InsertionSortFront(LessStr)
	l.RangeFront(func(v string) bool {
		fmt.Printf("%v\n", v)
		return true
	})
	fmt.Printf("%v\n", l.Size())
	// Output:
	// aaaaa
	// bbbbb
	// ccccc
	// ddddd
	// 4
}

func Example_list3() {
	l := New[string]()
	l.PushBack("ddddd")
	l.PushFront("ccccc")
	l.PushBack("aaaaa")
	l.PushFront("bbbbb")

	l.InsertionSortBack(LessStr)
	l.RangeFront(func(v string) bool {
		fmt.Printf("%v\n", v)
		return true
	})
	fmt.Printf("%v\n", l.Size())
	// Output:
	// ddddd
	// ccccc
	// bbbbb
	// aaaaa
	// 4
}

func Example_list4() {
	l := New[string]()
	l.PushBack("ddddd")
	l.PushFront("ccccc")
	l.PushBack("aaaaa")
	l.PushFront("bbbbb")

	l.InsertionSortFront(LessStr)
	l.RangeBack(func(v string) bool {
		fmt.Printf("%v\n", v)
		return true
	})
	fmt.Printf("%v\n", l.Size())
	// Output:
	// ddddd
	// ccccc
	// bbbbb
	// aaaaa
	// 4
}

func Example_list5() {
	l := New[string]()
	l.PushBack("ddddd")
	l.PushFront("ccccc")
	l.PushBack("aaaaa")
	l.PushFront("bbbbb")

	l.InsertionSortBack(LessStr)
	l.RangeBack(func(v string) bool {
		fmt.Printf("%v\n", v)
		return true
	})
	fmt.Printf("%v\n", l.Size())
	// Output:
	// aaaaa
	// bbbbb
	// ccccc
	// ddddd
	// 4
}

func Test1(t *testing.T) {

}
