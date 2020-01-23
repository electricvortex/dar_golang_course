package main

import (
	"fmt"
)

func main(){
	//var a *int
	//fmt.Println(a)
	//panic("Can not run")
	//*a = 6
	//b := 6
	//a = &b
	//fmt.Println(a)

	//var c interface{} = 7.0
	//switch c.(type) {
	//case int:
	//	fmt.Println("C is int type")
	//case string:
	//	fmt.Println("C is string type")
	//case float32:
	//	fmt.Println("C is float32 type")
	//case float64:
	//	fmt.Println("C is float64 type")
	//default:
	//	fmt.Println("Unknown type")
	//}
	//
	//m := make(map[string]int, 10)
	//fmt.Println(m)

	//arr := [6]int{1,2,3,4,5,6}
	////arr2 := [10]int{}
	//slice := arr[0:3]
	//slice = append(slice, 5, 5, 6, 6, 6)
	//fmt.Println(slice)
	//fmt.Println("length: ", len(slice))
	//fmt.Println("capacity: ", cap(slice))
	//slice = append(slice, 9,9,9,9,9)
	//fmt.Println(slice)
	//fmt.Println("length: ", len(slice))
	//fmt.Println("capacity: ", cap(slice))
	//fmt.Println("capacity of the array: ", cap(arr))
	//
	//// Variadic
	//slice = append([]int{1, 1, 1}, slice...)
	//fmt.Println(slice)

	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%v %v \n", i, IsOdd(i))
	//}
	//
	//l, k := Swap(1, 2)
	//fmt.Println(l, k)

	//arr := []int{1, 2, 3}
	//add(&arr, 7)
	//fmt.Println(arr)

	//d := DoubleX(2)
	//fmt.Println(d(1))
	//fmt.Println(d(3))

	//fmt.Println(Fibo(6))

	//dino := &Dino{}
	//PrintWalking(dino)

	//var a string
	//a = "5"
	//
	//b, err := strconv.Atoi(a)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(b)
	//fmt.Println(err.Error())

	//n1 := make(chan int)
	//
	//go printNumbers(n1, 0)
	//go printNumbers(n1, 10)
	//
	//for i:=0; i<20; i++ {
	//	fmt.Println(<-n1)
	//}


	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func IsOdd(n int) string {
	if n % 2 == 0 {
		return "is even"
	} else {
		return "is odd"
	}
}

//func Swap(a, b int) (int, int) {
//	return b, a
//}

func Swap(a, b int) (c, d int) {
	c = b
	d = a
	return
}

func add(a *[]int, b ...int) {
	*a = append(*a, b...)
	fmt.Println(a)
}

func DoubleX(x int) func(n int) int {
	l := x
	return func(n int) int {
		l = l * l + n
		return l
	}
}

func Fibo(x int) int {
	if x < 2 {
		return 1
	}
	return Fibo(x - 1) + Fibo(x - 2)
}

type DinoInt interface {
	Walk()
}

type Dino struct {
	Name string
}

func (d *Dino) Walk() {
	fmt.Println("walking...")
}

func PrintWalking(d DinoInt) {
	d.Walk()
}

func printNumbers(a chan <- int, n int) {
	for i := n; i < n+10; i++ {
		a <- i
	}
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
