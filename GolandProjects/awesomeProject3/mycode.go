package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("about to sleep in main()")
	time.Sleep(10 * 1e9)
	fmt.Println("at the end of main()")
}
func longWait() {
	fmt.Println("Beginning longWait")
	time.Sleep(5 * 1e9)
	fmt.Println("End of longWait")
}
func shortWait() {
	fmt.Println("Beginning of shortWait")
	time.Sleep(2 * 1e9)
	fmt.Println("End of shortWait")
}
