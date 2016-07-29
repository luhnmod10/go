// Package luhnmod10 is a fast and simple in-place implementation of the luhn check algorithm.
package luhnmod10

// Valid returns true if the number string is luhn valid, and false otherwise.
// The number passed to the function must contain only numeric characters otherwise behavior is undefined.
func Valid(number string) bool {
	var checksum int

	numberLen := len(number)
	for i := numberLen - 1; i >= 0; i -= 2 {
		n := number[i] - '0'
		checksum += int(n)
	}
	for i := numberLen - 2; i >= 0; i -= 2 {
		n := number[i] - '0'
		n *= 2
		if n > 9 {
			n -= 9
		}
		checksum += int(n)
	}

	return checksum%10 == 0
}
