package main

import (
	"fmt"
	"src/src/mathutils"
)

func main() {
	fmt.Println("Hello World!")
	floatArray := []float64{15.0, 30.0}
	fmt.Println(mathutils.Average(floatArray))
}
