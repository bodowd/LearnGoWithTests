package integers

import "fmt"

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func Add(x, y int) int {
	return x + y
}
