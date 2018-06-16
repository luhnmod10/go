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

func ExampleChecksum() {
	number := "4242424242424241"
	checksum := luhnmod10.Checksum(number)
	fmt.Println(number, "checksum:", checksum)

	// Output:
	// 4242424242424241 checksum: 79
}

func ExampleCheckDigit() {
	number := "424242424242424"
	checkDigit := luhnmod10.CheckDigit(number)
	fmt.Println(number, "check digit:", checkDigit)

	// Output:
	// 424242424242424 check digit: 2
}
