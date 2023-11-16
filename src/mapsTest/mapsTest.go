package mapsTest

import "fmt"

func TestMaps() {
	myMap := map[string]string{}
	myMap["fistName"] = "Lucas"
	myMap["lastName"] = "Reginatto"
	fmt.Println(myMap)
	myMapDeclared := map[string]string{
		"firstName": "John",
		"lastname":  "Doe",
	}
	fmt.Println(myMapDeclared)

}
