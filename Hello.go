package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"time"
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

type DivideError struct {
	divideDee int
	divideDer int
}

func (de *DivideError) Error() string {
	stringFormat := `Cannot proceed, the divider is zero.
    dividee: %d
	divider: 0
	`
	return fmt.Sprintf(stringFormat, de.divideDee)
}

func Divide(numerator int, denominator int) (result int, errorMsg string) {
	if denominator == 0 {
		errData := DivideError{
			divideDee: numerator,
			divideDer: denominator,
		}
		errorMsg = errData.Error()
		return
	} else {
		return numerator / denominator, ""
	}
}

func main() {
	//var x interface{}

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

	// switch x.(type) {
	// case nil:
	// 	println("x 是 nil 型")
	// 	break
	// }

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

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("All you need is real ", CalCircle(c1))

	p := Brother()
	fmt.Println("Is that yours:", p())
	fmt.Println("Is that yours:", p())

	var arr = []float32{1.0, 2.0, 3.0, 4.5}
	// this like foreach
	for index, value := range arr {
		fmt.Println("the value is ", index, value)
	}

	//leizhou := Person{ name="leizhou", age }

	v1 := []int{1, 2, 3}

	fmt.Println("know you", v1[1:])

	country := map[string]string{"China": "Beijing", "Japan": "Tokyo", "USA": "WashionDon D.C"}
	//country = make(map[string]string)

	//country["China"] = "Beijing"
	delete(country, "Japan")
	for _, item := range country {
		fmt.Println(item, country[item])
	}
	mon := new(GoldenMonkey)
	mon.name()

	if _, msg := Divide(3, 1); msg != "" {
		fmt.Println(msg)
	}

	go say("周")
	say("磊")

	s := []int{7, 2, 8, -9, 4, 0}
	channel := make(chan int)
	go sum(s[:len(s)/2], channel)
	go sum(s[len(s)/2:], channel)
	x, y := <-channel, <-channel
	fmt.Println(x, y, x+y)

	ch := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2
	//ch <- 3
	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	rand.Seed(100)
	fmt.Println("random :", rand.Intn(100))
	for index := 0; index < len(s); index++ {
		fmt.Println("use for ", s[index])
	}
	fmt.Println("math :", math.Pi)

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("To")
		break
	case today + 1:
		fmt.Println("one day later")
		break
	default:
		fmt.Println("too far")
	}

	defer fmt.Println("I appear later")
	fmt.Println("I'm showing right now")
	res, err := http.Get("https://kyfw.12306.cn/passport/captcha/captcha-image64?login_site=E&module=login&rand=sjrand&1553139749524&callback=jQuery19106227248862412864_1553139737622&_json_att=&_=1553139737624")
	if err != nil {
		fmt.Println("error happend", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	//var data ImageJson
	//json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", string(body))
}

type ImageJson struct {
	Image         string
	ResultCode    string
	ResultMessage string
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func CalCircle(a Circle) float64 {
	return 3.14 * a.radius * a.radius
}

func GoGo() (int, int) {
	return 1, 2
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

type Circle struct {
	radius float64
}

type Person struct {
	name   string
	age    int
	Height float32
}

type Monkey interface {
	kind()
	name()
}

type GoldenMonkey struct {
}

func (monkey GoldenMonkey) name() {
	fmt.Println("I'm a monkey")
}

func Brother() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func sum(arr []int, channel chan int) {
	sum := 0
	for index := 0; index < len(arr); index++ {
		sum += arr[index]
	}
	channel <- sum
}
