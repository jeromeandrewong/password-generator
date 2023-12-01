package password

import (
	"errors"
	"strings"
	"testing"
)

type MockInputGetter struct {
	inputs []string
	index  int
}

func (m *MockInputGetter) GetInput() (string, error) {
	if m.index >= len(m.inputs) {
		return "", errors.New("no more input")
	}
	val := m.inputs[m.index]
	m.index++
	return val, nil
}

func TestGetUserInput(t *testing.T) {
	// Stubbing user input for testing
	mockInput := "12\n3\n4\n2\n3\n"
	uiStub := &MockInputGetter{
		inputs: strings.Fields(mockInput),
		index:  0,
	}

	// Replace the original GetInput function with the stub for testing
	userInput = uiStub

	expectedReq := PasswordRequirements{
		Length:  12,
		Lower:   4,
		Upper:   3,
		Numbers: 2,
		Special: 3,
	}

	req, err := GetUserInput()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if req != expectedReq {
		t.Errorf("Expected %v, but got %v", expectedReq, req)
	}
}
