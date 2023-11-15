package loops

import "fmt"

func TestWhile() {
	a := 0
	for a < 4 {
		fmt.Println(a)
		a++
	}
}
