package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I'm a string"
	fmt.Printf("End with string? %v\n", strings.HasSuffix(s, "string"))
}
