package variables

// cause of the lowercase, the function is private
func exampleFunc(name string) string {
	return "Hello " + name
}

// cause of the uppercase, the function is public (accessible outside the package)
func ExampleFunc2(name string) string {
	return "Hello " + name
}

// by default, go passes by value, so the value of x is not modified
func ModifyValue(x int) int {
	x = 10
	return x
}

// by using a pointer, we can modify the value of x
func ModifyValue2(x *int) {
	*x = 10
}
