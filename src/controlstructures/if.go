package controlstructures

import "fmt"

func TestIf() {
	a := 5
	if a < 0 {
		fmt.Println("Your value is negative!")
	} else if a < 10 {
		fmt.Println("Your value is a sigle digit!")
	} else {
		fmt.Println("Your value has multiple digits!")
	}

}
