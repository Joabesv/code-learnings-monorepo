package main

import (
	"fmt"
	variables "v0/modules"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(variables.GetAge())
	x := 34
	fmt.Println("og x value", x)
	// wont modify the value of x
	variables.ModifyValue(x)
	fmt.Println("non-pointer", x)
	// will modify the value of x
	variables.ModifyValue2(&x)
	fmt.Println("pointer", x)
}
