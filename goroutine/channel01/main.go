package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func FirsthalfCal(s, max int, ch chan int) {
	for ; s <= max; s++ {
		ch <- s
		//fmt.Println(s)
	}
	defer close(ch)
}

func CheckChan(ch chan int, ch2 chan int) {
	_, ok := <-ch
	if ok == false {
		_, ok2 := <-ch2
		if ok2 == false {
			time.Sleep(3 * time.Second)
			os.Exit(0)
		}
	}
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go FirsthalfCal(2000, 4000, ch)
	go FirsthalfCal(0, 2000, ch2)

	for {
		select {
		case i, ok := <-ch:
			if ok == false {
				//CheckChan(ch, ch2)
				//time.Sleep(1 * time.Second)
			} else {
				fmt.Println(i)
			}
		case m, ok := <-ch2:
			if ok == false {
				//CheckChan(ch, ch2)
				//time.Sleep(1 * time.Second)
			} else {
				fmt.Println(m)
			}
		default:
			fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())
			time.Sleep(5 * time.Second)
		}

	}
}
