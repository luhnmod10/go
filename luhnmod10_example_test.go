package luhnmod10_test

import (
	"fmt"

	"github.com/luhnmod10/go"
)

func ExampleValid_true() {
	number := "4242424242424242"
	valid := luhnmod10.Valid(number)
	fmt.Println(number, "is valid:", valid)

	// Output:
	// 4242424242424242 is valid: true
}

func ExampleValid_false() {
	number := "4242424242424241"
	valid := luhnmod10.Valid(number)
	fmt.Println(number, "is valid:", valid)

	// Output:
	// 4242424242424241 is valid: false
}
