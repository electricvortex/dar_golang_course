package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	Atomic()

}

func WaitGroupAndDefer() {
	defer fmt.Println("One")
	defer fmt.Println("Second")

	channel := make(chan int, 200)

	w := sync.WaitGroup{}

	for i := 1; i<=2; i++ {
		w.Add(1)
		go func(i int, temp chan<- int, group *sync.WaitGroup) {
			for j := (i - 1) * 10; j < i * 10; j++ {
				temp <- j
			}
			defer group.Done()
		}(i, channel, &w)
	}

	w.Wait()
	close(channel)

	for v := range channel {
		fmt.Println(v)
	}
}

func Atomic() {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}