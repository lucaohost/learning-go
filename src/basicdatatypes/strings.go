package basicdatatypes

import (
	"fmt"
	"strings"
)

func TestStrings() {
	s := "I'm a string"
	fmt.Printf("End with string? %v\n", strings.HasSuffix(s, "string"))
}
