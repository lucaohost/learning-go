package main

import (
	"fmt"
	"mathutils/src/mathutils"
)

func main() {
	fmt.Println("Hello World!")
	floatArray := []float64{15.0, 30.0}
	fmt.Println(mathutils.Average(floatArray))
}
