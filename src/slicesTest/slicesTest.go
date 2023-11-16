package slicesTest

import "fmt"

func TestSlices() {
	myArray := []int{1, 2, 3, 4, 5}
	myArray = append(myArray, 6)
	fmt.Println(myArray)
	myArray[2] = 60
	fmt.Println(myArray)
}
