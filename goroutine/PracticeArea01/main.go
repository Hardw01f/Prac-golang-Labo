package main

import (
	"fmt"
)

/*func sub() {
	for {
		fmt.Println("This is sub loop")
	}
}
*/

func main() {
	go func(){
			for {
					fmt.Println("This is sub loop")
			}
	}()

	for {
		fmt.Println("This is MAIN loop")
	}
}
