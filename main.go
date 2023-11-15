package main

import (
	"fmt"

	"github.com/lucaohost/learning-go/src/basicdatatypes"
	"github.com/lucaohost/learning-go/src/controlstructures"
	"github.com/lucaohost/learning-go/src/loops"
	"github.com/lucaohost/learning-go/src/mathutils"
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
}
