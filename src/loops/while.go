package loops

import "fmt"

func testWhile() {
	a := 0
	for a < 4 {
		fmt.Println(a)
		a++
	}
}
