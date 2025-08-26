package methods

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

// method example, it is a function with a receiver
// receiver is a pointer to the Employee struct
// the method is defined outside the struct
// this is a value receiver, so the method is called on a copy of the struct
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s %d", e.name, e.currency, e.salary)
}

func DisplayInfo() {
	emp := Employee{
		name:     "John",
		salary:   1000,
		currency: "USD",
	}

	// later the method is called like a function of the struct
	emp.displaySalary()

	fmt.Println("Before changing name")
	emp.changeName("Joabe")
	emp.displaySalary()

	fmt.Println("After changing name")
	emp.displaySalary()

}

// this is a pointer receiver, so the method is called on the original struct
func (e *Employee) changeName(newName string) {
	e.name = newName
}
