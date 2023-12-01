package password

import (
	"fmt"
	"strings"
)

var (
	lower    = "abcdefghijklmnopqrstuvwxyz"
	upper    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers  = "0123456789"
	special  = "!@#$%^&*()_+={}[]|<>/?~`"
	allChars = lower + upper + numbers + special
)

func GeneratePassword() {
	password, err := CreatePassword()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(password)
}

func CreatePassword() (string, error) {
	passReq, err := GetUserInput()
	if err != nil {
		return "", err
	}

	var password strings.Builder

	for i := 0; i < passReq.Lower; i++ {
		n, err := RandomInt64(len(lower))
		if err != nil {
			return "", err
		}
		password.WriteString(string(lower[n]))
	}

	for i := 0; i < passReq.Upper; i++ {
		n, err := RandomInt64(len(upper))
		if err != nil {
			return "", err
		}
		password.WriteString(string(upper[n]))
	}

	for i := 0; i < passReq.Numbers; i++ {
		n, err := RandomInt64(len(numbers))
		if err != nil {
			return "", err
		}
		password.WriteString(string(numbers[n]))
	}

	for i := 0; i < passReq.Special; i++ {
		n, err := RandomInt64(len(special))
		if err != nil {
			return "", err
		}
		password.WriteString(string(special[n]))
	}

	remaining := passReq.Length - passReq.Lower - passReq.Upper - passReq.Numbers - passReq.Special

	for i := 0; i < remaining; i++ {
		n, err := RandomInt64(len(allChars))
		if err != nil {
			return "", err
		}
		password.WriteString(string(allChars[n]))
	}

	pass, err := ScramblePassword(password.String())
	if err != nil {
		return "", err
	}

	return pass, nil
}
