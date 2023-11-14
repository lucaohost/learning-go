package basicdatatypes

import (
	"fmt"
	"strings"
)

func testStrings() {
	s := "I'm a string"
	fmt.Printf("End with string? %v\n", strings.HasSuffix(s, "string"))
}
