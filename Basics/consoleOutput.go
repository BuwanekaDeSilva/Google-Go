package main

import "fmt"

func main() {

	str1 := "The quick brown fox"
	str2 := "jumps over"
	str3 := "the lazy dog"
	aNumber := 50
	isTrue := true

	stringLenth, _ := fmt.Println(str1, str2, str3) // _ auto initialize varible to the next output statement

	// if err == nil {
	fmt.Println("String Length: ", stringLenth)
	// }

	fmt.Printf("Value of a Number %v\n", aNumber)
	fmt.Printf("Value of a boolean %v\n", isTrue)
	fmt.Printf("Value of a float %.2f\n", float64(aNumber)) //conversion to float

	fmt.Printf("Data Types : %T,%T,%T,%T and %T \n", str1, str2, str3, aNumber, isTrue) //output data types

	thisString := fmt.Sprintf("Data Types as var : %T,%T,%T,%T and %T", str1, str2, str3, aNumber, isTrue)

	fmt.Print(thisString) //output as strings
}
