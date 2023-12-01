package main

import (
	"bufio"
	cRand "crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	lower    = "abcdefghijklmnopqrstuvwxyz"
	upper    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers  = "0123456789"
	special  = "!@#$%^&*()_+={}[]|<>/?~`"
	allChars = lower + upper + numbers + special
)

type PasswordRequirements struct {
	Length  int
	Lower   int
	Upper   int
	Numbers int
	Special int
}

func main() {
	password, err := createPassword()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)

}

func createPassword() (string, error) {
	passReq, err := getUserInput()
	if err != nil {
		return "", err
	}

	var password strings.Builder

	for i := 0; i < passReq.Lower; i++ {
		n, err := randomInt64(len(lower))
		if err != nil {
			return "", err
		}
		password.WriteString(string(lower[n]))
	}

	for i := 0; i < passReq.Upper; i++ {
		n, err := randomInt64(len(upper))
		if err != nil {
			return "", err
		}
		password.WriteString(string(upper[n]))
	}

	for i := 0; i < passReq.Numbers; i++ {
		n, err := randomInt64(len(numbers))
		if err != nil {
			return "", err
		}
		password.WriteString(string(numbers[n]))
	}

	for i := 0; i < passReq.Special; i++ {
		n, err := randomInt64(len(special))
		if err != nil {
			return "", err
		}
		password.WriteString(string(special[n]))
	}

	remaining := passReq.Length - passReq.Lower - passReq.Upper - passReq.Numbers - passReq.Special

	for i := 0; i < remaining; i++ {
		n, err := randomInt64(len(allChars))
		if err != nil {
			return "", err
		}
		password.WriteString(string(allChars[n]))
	}

	pass, err := scramblePassword(password.String())
	if err != nil {
		return "", err
	}

	return pass, nil
}

func randomInt64(len int) (int64, error) {
	n, err := cRand.Int(cRand.Reader, big.NewInt(int64(len)))
	return n.Int64(), err
}

func scramblePassword(s string) (string, error) {
	r := []rune(s)
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})

	return string(r), nil
}

func getUserInput() (PasswordRequirements, error) {

	p := PasswordRequirements{}
	reqCount := 0

	fmt.Print("how long should the password be? ")
	str, err := getInput()
	if err != nil {
		return PasswordRequirements{}, err
	}
	p.Length, err = strconv.Atoi(str)
	if err != nil {
		return PasswordRequirements{}, errors.New("please enter a valid number")
	}

	fmt.Print("how many uppercase letters? ")
	str, err = getInput()
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

	fmt.Print("how many lowercase letters? ")
	str, err = getInput()
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

	fmt.Print("how many numbers should there be? ")
	str, err = getInput()
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

	fmt.Print("how many special characters should there be? ")
	str, err = getInput()
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

func getInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	return scanner.Text(), nil
}
