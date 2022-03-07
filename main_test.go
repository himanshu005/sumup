package main

import "testing"

//Function #1: Validity of the card number

func TestIsValidCreditCardNumber(t *testing.T) {
	var tests = []struct {
		str string
		exp bool
	}{
		{"5237 2516 2477 8133 ", true},
		{"5237 2516 2477 8133 1", false}, //negative test case
	}
	for _, e := range tests {
		res := isValidCreditCardNumber(e.str)
		if res != e.exp {
			t.Errorf("%v is Valid Credit Card Number \n actual: %v \n expected: %v",
				e.str, res, e.exp)
		}
	}
}

// Function #2: Known/supported card schemes

func TestGetSupportedCardSchemes(t *testing.T) {
	var tests = []struct {
		str string
		exp string
	}{
		{"378282246310005", "American Express"},
		{"3530111333300000", "JCB"},
		{"6759649826438453", "Maestro"},
		{"4012888888881881", "Visa"},
		{"5105105105105100", "MasterCard"},
		{"5237 2516 2477 8133 1", "none"}, //negative test case
	}
	for _, e := range tests {
		res := getSupportedCardSchemes(e.str)
		if res != e.exp {
			t.Errorf("%v Get Supported Card Schemes\n actual: %v \n expected: %v",
				e.str, res, e.exp)
		}
	}
}
