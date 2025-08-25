package variables

import "fmt"

func GetConstants() {
	const n = "Sam"
	var name = n
	fmt.Printf("type %T value %v", name, name)
}
