package main

import (
		"fmt"
)

func Hello(str string,ch chan string){
		//fmt.Println("I am ",str)
		ch <- "I am " + str
		return ch
}

func main(){
		ch := make(chan string)
		c := Hello("Ironman",ch)
		fmt.Println(c)
}
