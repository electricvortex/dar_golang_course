package main

import "fmt"

const a float32 = 123

func main(){
	const b float32 = 124
	a := 5
	if a > 6 {
		fmt.Println("Less")
	} else {
		fmt.Println("OK")
	}


	a1 := [10]int{}
	for i:=0; i < 10; i++ {
		a1[i] = i
	}
	b1 := a1[1:4]
	b1[2] = 9
	fmt.Println(b)
	fmt.Println(a)

	for i, _ := range a1 {
		fmt.Println(a1[i])
	}

	b2 := "hello world"
	for i, v := range b2 {
		fmt.Printf("Index is : %v \n", i)
		fmt.Printf("Values is: %v \n", string(v))
	}

}