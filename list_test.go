package list

import "fmt"
import "testing"

func ExampleList1() {
	l := New()
	l.PushBack("ddddd")
	l.PushFront("ccccc")
	l.PushBack("aaaaa")
	l.PushFront("bbbbb")
	
	for it := l.Front(); it != l.End(); it = it.Next() {
		fmt.Printf("%v\n", it.Value)
	}
	fmt.Printf("%v\n", l.Size())
/* Output:
bbbbb
ccccc
ddddd
aaaaa
4
*/
}

type MyLess_t struct {}

func (MyLess_t) Less(a interface{}, b interface{}) bool {
	return a.(string) < b.(string)
}

func ExampleList2() {
	l := New()
	l.PushBack("ddddd")
	l.PushFront("ccccc")
	l.PushBack("aaaaa")
	l.PushFront("bbbbb")
	
	l.InsertionSortFront(MyLess_t{})
	for it := l.Front(); it != l.End(); it = it.Next() {
		fmt.Printf("%v\n", it.Value)
	}
	fmt.Printf("%v\n", l.Size())
/* Output:
aaaaa
bbbbb
ccccc
ddddd
4
*/
}

func ExampleList3() {
	l := New()
	l.PushBack("ddddd")
	l.PushFront("ccccc")
	l.PushBack("aaaaa")
	l.PushFront("bbbbb")
	
	l.InsertionSortBack(MyLess_t{})
	for it := l.Front(); it != l.End(); it = it.Next() {
		fmt.Printf("%v\n", it.Value)
	}
	fmt.Printf("%v\n", l.Size())
/* Output:
ddddd
ccccc
bbbbb
aaaaa
4
*/
}

func Test1(t * testing.T) {

}
