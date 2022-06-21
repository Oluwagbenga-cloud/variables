//basic map study
package main

import "fmt"

func main() {
	var app map[string]int
	fmt.Println(app)
	if app == nil {
		fmt.Println("app variable is nil")
	}
}
