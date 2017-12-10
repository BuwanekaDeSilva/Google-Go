package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "An implicitly typed string"
	fmt.Print("str1: %V:%T\n", str1, str1)

	str2 := "An implicitly typed string"
	fmt.Print("str2: %V:%T\n", str2, str2)

	fmt.Print(strings.ToUpper(str2))

	lvalue := "hello"
	uvalue := "HELLO"
	fmt.Print("Equals?\n", (lvalue == uvalue))
	fmt.Print("Equals Non-case snstive?\n", strings.EqualFold(lvalue, uvalue))
}
