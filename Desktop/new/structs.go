package main

import "fmt"

func main() {
	type UserInfo struct {
		Name string
		Age  int
		Id   int
	}
	var info []UserInfo
	info = append(info, UserInfo{Name: "Gbenga", Age: 33, Id: 1})
	info = append(info, UserInfo{Name: "Kayla", Age: 30, Id: 2})
	info = append(info, UserInfo{Name: "Iyinoluwa Adrian", Age: 1, Id: 3})
	fmt.Println(info)
	//apps := []string{"Facebook", "Instagram", "Whatsapp", "Discord", "Reddit"}
	//
	//for i, app := range apps {
	//	fmt.Println(i, app)
	//}

}
