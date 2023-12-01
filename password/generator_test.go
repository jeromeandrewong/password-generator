package password

import (
	"testing"
)

func TestCreatePassword(t *testing.T) {

	// Test case 1: Test if the password length matches the requirements
	passReq := PasswordRequirements{
		Length:  12,
		Lower:   3,
		Upper:   3,
		Numbers: 3,
		Special: 3,
	}

	password, err := CreatePassword(passReq)
	if err != nil {
		t.Errorf("Error creating password: %s", err)
	}

	if len(password) != passReq.Length {
		t.Errorf("Expected password length %d, got %d", passReq.Length, len(password))
	}

	// Test case 2: Password length exceeding the specified length
	passReq = PasswordRequirements{
		Length:  15,
		Lower:   5,
		Upper:   5,
		Numbers: 5,
		Special: 5,
	}

	_, err = CreatePassword(passReq)
	if err == nil {
		t.Errorf("Expected an error due to requirements exceeding password length")
	}

}
