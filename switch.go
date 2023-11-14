package main

import "fmt"

func main() {
	a := 10
	switch a {
	case 1:
		fmt.Println("The value is 1!")
	case 10:
		fmt.Println("The value is 10!")
	default:
		fmt.Println("None of these!")
	}
}
