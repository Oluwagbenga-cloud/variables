//basic map study
package main

import "fmt"

func main() {
	//var app = make(map[string]int)
	//app["dividend"] = 21
	//fmt.Println(app)
	app := make(map[int]string)
	app[21] = "Facebook"
	app[22] = "Whatsapp"
	app[23] = "LinkedIn"
	app[24] = "Instagram"
	app[25] = "Reddit"
	fmt.Println(app)
}
