package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum :", intSum)

	f1, f2, f3 := 23.5, 64.5, 76.3
	fSum := f1 + f2 + f3
	fmt.Println("float sum :", fSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(23.5)
	b2.SetFloat64(64.5)
	b3.SetFloat64(76.3)
	bigSum.Add(&b1, &b2).Add(&b3, &bigSum)
	fmt.Printf("Bigsum = %.10g\n", &bigSum)

	CircleRadius := 15.5
	circumfernece := CircleRadius * math.Pi
	fmt.Printf("Circumference= %.2F\n", circumfernece)

}
