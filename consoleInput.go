package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var s string
	// fmt.Scanln(&s)
	// fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter some text: ")

	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter some Number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	fmt.Println(str)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The output value", f)
	}

}
