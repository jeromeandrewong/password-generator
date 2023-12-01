package password

import "testing"

func TestScramblePassword(t *testing.T) {

	// Test case 1: Test if the scrambled password length matches the input
	password := "TestPassword123!"
	scrambledPassword, err := ScramblePassword(password)
	if err != nil {
		t.Errorf("Error scrambling password: %s", err)
	}

	if len(scrambledPassword) != len(password) {
		t.Errorf("Expected scrambled password length %d, got %d", len(password), len(scrambledPassword))
	}

	// Test case 2: Test when the input is an empty string
	emptyPassword := ""
	emptyScrambledPassword, err := ScramblePassword(emptyPassword)
	if err != nil {
		t.Errorf("Error scrambling empty password: %s", err)
	}

	if emptyScrambledPassword != "" {
		t.Errorf("Expected empty string for scrambled empty password")
	}
}

func TestRandomInt64(t *testing.T) {
	testLength := 10
	randomNumber, err := RandomInt64(testLength)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	maxValue := int64(testLength)
	if randomNumber < 0 || randomNumber >= maxValue {
		t.Errorf("Random number out of range: %d", randomNumber)
	}
}
