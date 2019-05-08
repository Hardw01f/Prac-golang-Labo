package main

import (
	"fmt"
	"sync"
)

func Return(wg *sync.WaitGroup) {
	fmt.Println("I am Ironman")
	wg.Done()
}

func main() {
	//ch := make(chan int)
	//ch2 := make(chan int)
	wg := new(sync.WaitGroup)

	for i := 0; i <= 200; i++ {

		if i%2 == 0 {
			wg.Add(1)
			go Return(wg)
		} else if i%2 == 1 {
			wg.Add(1)
			go Return(wg)
		}
	}

	wg.Wait()
}
