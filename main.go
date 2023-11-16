package main

import (
	"fmt"

	"github.com/lucaohost/learning-go/src/basicdatatypes"
	"github.com/lucaohost/learning-go/src/controlstructures"
	"github.com/lucaohost/learning-go/src/errorsTesting"
	"github.com/lucaohost/learning-go/src/loops"
	"github.com/lucaohost/learning-go/src/mathutils"
	"github.com/lucaohost/learning-go/src/pointers"
)

func main() {
	fmt.Println("Hello World!")
	floatArray := []float64{15.0, 30.0}
	fmt.Println(mathutils.Average(floatArray))
	basicdatatypes.TestBooleans()
	basicdatatypes.TestNumbers()
	basicdatatypes.TestStrings()
	controlstructures.TestIf()
	controlstructures.TestSwitch()
	loops.TestInfiniteLoop()
	loops.TestLoop()
	loops.TestRange()
	loops.TestWhile()
	errorsTesting.TestErrors()
	println("=============================================\n")
	pointers.TestPointers()
	println("\n=============================================")
}
