package pointer

import "fmt"

func PointerExample() {
	b := 255
	var a *int = &b

	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
}
