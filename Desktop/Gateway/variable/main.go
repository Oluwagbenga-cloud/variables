package main

import "fmt"

const TimberAnthony string = "gibberish"

func main(){

	var username string = "Oluwagbenga"
fmt.Println(username)
fmt.Printf("Variable is of type: %T \n", username)

var isLoggedIn bool = true
fmt.Println(isLoggedIn)
fmt.Printf("Variable is of type: %T \n", isLoggedIn)
 
var smallVal uint8 = 255
fmt.Println(smallVal)
fmt.Printf("Variable is of type: %T \n", smallVal)

var smallFloat float32 = 255.89 
fmt.Println(smallFloat)
fmt.Printf("Variable is of type: %T \n", smallFloat)

var website = "visitoja.com"
fmt.Println(website)

fmt.Println(TimberAnthony)
fmt.Printf("Variable is of the type: %T \n", TimberAnthony)
}


