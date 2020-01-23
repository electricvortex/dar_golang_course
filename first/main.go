package main

import (
	"fmt"
)

const a float32 = 123

func main(){
	//const b float32 = 124
	//a := 5
	//if a > 6 && a < 10 {
	//	fmt.Println("Less")
	//} else {
	//	fmt.Println("OK")
	//}
	//
	//a1 := [10]int{}
	//for i:=0; i < 10; i++ {
	//	a1[i] = i
	//}
	//b1 := a1[1:4]
	//b1[2] = 9
	//fmt.Println(b)
	//fmt.Println(a)
	//
	//for i, _ := range a1 {
	//	fmt.Println(a1[i])
	//}
	//
	//b2 := "hello world"
	//for i, v := range b2 {
	//	fmt.Printf("Index is : %v \n", i)
	//	fmt.Printf("Values is: %v \n", string(v))
	//}
	//
	//first := 10
	//for first > 10 {
	//	first--
	//}
	//first += 10
	//
	//var l bool = true
	//

	//m, err := strconv.Atoi("c")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(m)
	//
	//strconv.Itoa(5)
	//strconv.Atoi("6")


	var l interface{} = 5.5

	switch l.(type) {
	case int:
		fmt.Println("Values is integer")
	case float32:
		fmt.Println("Values if float32")
	case float64:
		fmt.Println("Values is float64")
	default:
		fmt.Println("Unknown type")
	}



}