package basicdatatypes

import "fmt"

func TestNumbers() {
	i := 100
	var j int = 1234
	fmt.Printf("%v + %v = %v\n", i, j, i+j)
	f := 1.5
	fmt.Printf("f = %v\n", f)
}
