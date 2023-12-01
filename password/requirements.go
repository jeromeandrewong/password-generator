package password

import (
	"errors"
	"fmt"
	"strconv"
)

type PasswordRequirements struct {
	Length  int
	Lower   int
	Upper   int
	Numbers int
	Special int
}

var userInput InputGetter = &RealInputGetter{}

func GetUserInput() (PasswordRequirements, error) {
	p := PasswordRequirements{}
	reqCount := 0

	fmt.Print("How long should the password be? ")
	str, err := userInput.GetInput()
	if err != nil {
		return PasswordRequirements{}, err
	}
	p.Length, err = strconv.Atoi(str)
	if err != nil {
		return PasswordRequirements{}, errors.New("please enter a valid number")
	}

	fmt.Print("Minimum number of uppercase letters? ")
	str, err = userInput.GetInput()
	if err != nil {
		return PasswordRequirements{}, err
	}
	p.Upper, err = strconv.Atoi(str)
	if err != nil {
		return PasswordRequirements{}, errors.New("please enter a valid number")
	}
	reqCount += p.Upper
	if reqCount > p.Length {
		return PasswordRequirements{}, errors.New("please enter a valid number")
	}

	fmt.Print("Minimum number of lowercase letters? ")
	str, err = userInput.GetInput()
	if err != nil {
		return PasswordRequirements{}, err
	}
	p.Lower, err = strconv.Atoi(str)
	if err != nil {
		return PasswordRequirements{}, errors.New("please enter a valid number")
	}
	reqCount += p.Lower
	if reqCount > p.Length {
		return PasswordRequirements{}, errors.New("exceeded password length")
	}

	fmt.Print("Minimum number of numbers? ")
	str, err = userInput.GetInput()
	if err != nil {
		return PasswordRequirements{}, err
	}
	p.Numbers, err = strconv.Atoi(str)
	if err != nil {
		return PasswordRequirements{}, errors.New("please enter a valid number")
	}
	reqCount += p.Numbers
	if reqCount > p.Length {
		return PasswordRequirements{}, errors.New("exceeded password length")
	}

	fmt.Print("Minimum number of special characters? ")
	str, err = userInput.GetInput()
	if err != nil {
		return PasswordRequirements{}, err
	}
	p.Special, err = strconv.Atoi(str)
	if err != nil {
		return PasswordRequirements{}, errors.New("please enter a valid number")
	}
	reqCount += p.Special
	if reqCount > p.Length {
		return PasswordRequirements{}, errors.New("exceeded password length")
	}

	return p, nil
}
