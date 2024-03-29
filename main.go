package main

import (
	"fmt"

	"github.com/lucaohost/learning-go/src/basicdatatypes"
	"github.com/lucaohost/learning-go/src/controlstructures"
	"github.com/lucaohost/learning-go/src/errorsTesting"
	"github.com/lucaohost/learning-go/src/interfacetest"
	"github.com/lucaohost/learning-go/src/loops"
	"github.com/lucaohost/learning-go/src/mapsTest"
	"github.com/lucaohost/learning-go/src/mathutils"
	"github.com/lucaohost/learning-go/src/methodstest"
	"github.com/lucaohost/learning-go/src/pointers"
	"github.com/lucaohost/learning-go/src/slicesTest"
	"github.com/lucaohost/learning-go/src/structures"
	"golang.org/x/sys/unix"
	"unicode/utf8"
)

type Person = structures.Person
type Phone = structures.Phone

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
	pointers.TestPointers()
	p := Person{
		FirstName: "Lucas",
		Lastname:  "Reginatto",
		Age:       25,
		Phone: Phone{
			Areacode: "55",
			Prefix:   "11",
			Suffix:   "999817226",
		},
	}
	fmt.Println(p)
	coordinate := struct {
		X int
		Y int
	}{
		X: 10,
		Y: 20,
	}
	fmt.Println(coordinate)
	slicesTest.TestSlices()
	mapsTest.TestMaps()
	p2 := methodstest.NewPerson("Lucas", "Reginatto", 1990, 10, 10)
	fmt.Println(p2.SayHello())
	fmt.Println(p2.GetAge())
	circle := &interfacetest.Circle{
		Radius: 10,
	}
	retangle := &interfacetest.Retangle{
		Width:  10,
		Height: 10,
	}
	PrintArea(circle)
	PrintArea(retangle)
	fmt.Println(GetUserId())
	println("=============================================\n")
	testRunes()
	println("\n=============================================")
}

func PrintArea(shape interfacetest.Shape) {
	fmt.Print(shape.Area(), "\n")
}

func GetUserId() int {
	return unix.Getuid()
}

func testRunes() {
	minhaString := "résumé"
    numeroDeCaracteres := utf8.RuneCountInString(minhaString)
	fmt.Printf("Número de bytes na string: %d\n", len(minhaString))
    fmt.Printf("Número de caracteres na string: %d\n", numeroDeCaracteres)
}
