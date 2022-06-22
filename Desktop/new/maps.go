//basic map study
package main

import "fmt"

func main() {
	//initializing the make function
	//var app = make(map[string]int)
	//app["dividend"] = 21
	//fmt.Println(app)

	//Initializing key values in a map
	//app := make(map[int]string)
	//app[21] = "Facebook"
	//app[22] = "Whatsapp"
	//app[23] = "LinkedIn"
	//app[24] = "Instagram"
	//app[25] = "Reddit"
	//fmt.Println(app)

	//initializing maps using map literal
	//var app = map[string]int{
	//	"k": 19,
	//	"l": 20,
	//	"m": 21,
	//}
	//fmt.Println(app)

	////deleting key from maps.
	//countries := make(map[int]string)
	//countries[1] = "Rwanda"
	//countries[2] = "Namibia"
	//countries[3] = "Kenya"
	//countries[4] = "Botswana"
	//countries[5] = "SouthAfrica"
	//countries[6] = "Tanzania"
	//fmt.Println(countries)
	//delete(countries, 5)

	age := map[string]int{

		"Nala":   29,
		"Kayla":  30,
		"Nadine": 30,
		"Busayo": 31,
	}
	BusayoAge, BusayoOk := age["Busayo"]
	NalaAge, NalaOk := age["Nala"]

	fmt.Println(BusayoAge, BusayoOk)
	fmt.Println(NalaAge, NalaOk)
}
