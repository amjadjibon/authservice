package utils

import (
	"testing"
)

func TestValidEmail(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"foo@gmail.com", true},
		{"Gopher <from@example.com>", true},
		{"example", false},
	}

	for _, test := range tests {
		addr, valid := ValidEmail(test.email)
		if valid != test.valid {
			t.Errorf("ValidEmail(%q) = %q, %v, want %v", test.email, addr, valid, test.valid)
		}
	}
}
