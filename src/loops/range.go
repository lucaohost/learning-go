package loops

import "fmt"

func TestRange() {
	for _, i := range []int{1, 2, 3, 4} {
		fmt.Println(i)
	}
}
