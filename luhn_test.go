package luhn

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
