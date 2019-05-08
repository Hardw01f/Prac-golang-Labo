package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")
	go Greet(1)
	go Greet(2)
	go Greet(3)
	time.Sleep(time.Second) //←追加したらちゃんと動く
	fmt.Println("main end")
}

func Greet(i int) {
	fmt.Println("hello", i)
}
