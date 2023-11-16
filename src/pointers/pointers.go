package pointers

import "fmt"

func TestPointers() {
	myString := "My String"
	memoryAddress := &myString
	fmt.Println(myString)
	fmt.Println(memoryAddress)
	fmt.Println(*memoryAddress)
}
