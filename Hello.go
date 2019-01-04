package main

import (
	"fmt"
	"unsafe"
)

var a int64

var (
	name string
	age  int
)

const DOU int = 16

const (
	Num1 = "NUM1"
	Num2 = unsafe.Sizeof(Num1)
	Num3 = len(Num1)
)

const (
	Go1 = iota
	Go2
	Go3
)

const (
	Test1 = 1 << iota
	Test2 = 3 << iota
	Test3
	Test4
)

func main() {
	var x interface{}

	b, c, d := 'a', 'b', 'c'
	_, f := 'd', 'e'
	fmt.Println(name, age, b, c, d, f)
	fmt.Println("Hello World")
	println(DOU)
	println(Num1, Num2, Num3)
	println(Go1, Go2, Go3)

	if Test1 > 1 && Test2 < 3 {
		println("是大于1 的")
	} else {
		println("是小于等于 1 的")
	}

	println(Test1, Test2, Test3, Test4)

	switch x.(type) {
	case nil:
		println("x 是 nil 型")
		break
	}

	// var c1, c2, c3 chan int
	// var i1, i2 int
	// select {
	// case i1 = <-c1:
	// 	println(i1)
	// case c2 <- i2:
	// 	print(i2)
	// case i3, ok := (<-c3):
	// 	if ok {

	// 	} else {
	// 		println(i3)
	// 	}
	// }

	var chg1, chg2 = 1, 2
	swap(&chg1, &chg2)
	fmt.Printf("after change %d, %d", chg1, chg2)
}

func GoGo() (int, int) {
	return 1, 2
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
