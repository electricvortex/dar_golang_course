package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	for i:=0; i<10; i++ {
		go func(w *sync.WaitGroup) {
			fmt.Println("hello ", i)
			w.Done()
		}(wg)
	}
	wg.Wait()
}
