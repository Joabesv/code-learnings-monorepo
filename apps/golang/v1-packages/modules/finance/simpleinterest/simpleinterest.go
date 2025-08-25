package simpleinterest

import "fmt"

func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}

func ListStuff() []int {
	a := [...]int{1, 2, 3}
	fmt.Println(a)
	return a[:]
}

func init() {
	fmt.Println("Simple interest module initialized")
}
