package main

import (
	"bytes"
	"fmt"
)

func main() {
	var chars [7]string // array of 7 elements
	chars[0] = "R"
	chars[1] = "A"
	chars[2] = "C"
	chars[3] = "E"
	chars[4] = "C"
	chars[5] = "A"
	chars[6] = "R"

	fmt.Println(":: This is an example ::") //output
	fmt.Println("************************") //output

	var buffers bytes.Buffer
	for i := 0; i < len(chars); i++ {
		buffers.WriteString(chars[i])
	}
	for i := len(chars) - 2; i >= 0; i-- {
		buffers.WriteString(chars[i])
	}

	fmt.Println(buffers.String()) //output as string

}
