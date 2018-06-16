package luhnmod10

import "testing"

func TestValid(t *testing.T) {
	tests := []struct {
		CardNumber    string
		ExpectedValid bool
	}{
		{"0", true},
		{"00", true},
		{"18", true},
		{"0000000000000000", true},
		{"4242424242424240", false},
		{"4242424242424241", false},
		{"4242424242424242", true},
		{"4242424242424243", false},
		{"4242424242424244", false},
		{"4242424242424245", false},
		{"4242424242424246", false},
		{"4242424242424247", false},
		{"4242424242424248", false},
		{"4242424242424249", false},
		{"42424242424242426", true},
		{"424242424242424267", true},
		{"4242424242424242675", true},
		{"000000018", true},
		{"99999999999999999999", true},
		{"99999999999999999999999999999999999999999999999999999999999999999997", true},
	}

	for _, test := range tests {
		valid := Valid(test.CardNumber)
		if valid == test.ExpectedValid {
			t.Logf("Valid(%#v) = %#v", test.CardNumber, valid)
		} else {
			t.Errorf("Valid(%#v) = %#v, expected %#v", test.CardNumber, valid, test.ExpectedValid)
		}
	}
}

func TestChecksum(t *testing.T) {
	tests := []struct {
		CardNumber       string
		ExpectedChecksum int
	}{
		{"0", 0},
		{"00", 0},
		{"18", 10},
		{"0000000000000000", 0},
		{"4242424242424240", 78},
		{"4242424242424241", 79},
		{"4242424242424242", 80},
		{"4242424242424243", 81},
		{"4242424242424244", 82},
		{"4242424242424245", 83},
		{"4242424242424246", 84},
		{"4242424242424247", 85},
		{"4242424242424248", 86},
		{"4242424242424249", 87},
		{"42424242424242493", 72},
		{"42424242424242426", 70},
		{"424242424242424267", 90},
		{"4242424242424242675", 80},
		{"000000018", 10},
		{"99999999999999999999", 180},
		{"99999999999999999999999999999999999999999999999999999999999999999997", 610},
	}

	for _, test := range tests {
		checksum := Checksum(test.CardNumber)
		if checksum == test.ExpectedChecksum {
			t.Logf("Checksum(%#v) = %#v", test.CardNumber, checksum)
		} else {
			t.Errorf("Checksum(%#v) = %#v, expected %#v", test.CardNumber, checksum, test.ExpectedChecksum)
		}
	}
}

func TestCheckDigit(t *testing.T) {
	tests := []struct {
		CardNumber         string
		ExpectedCheckDigit int
	}{
		{"00", 0},
		{"000", 0},
		{"424242424242424", 2},
	}

	for _, test := range tests {
		checkDigit := CheckDigit(test.CardNumber)
		if checkDigit == test.ExpectedCheckDigit {
			t.Logf("CheckDigit(%#v) = %#v", test.CardNumber, checkDigit)
		} else {
			t.Errorf("CheckDigit(%#v) = %#v, expected %#v", test.CardNumber, checkDigit, test.ExpectedCheckDigit)
		}
	}
}
