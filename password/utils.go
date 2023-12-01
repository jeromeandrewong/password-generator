package password

import (
	"bufio"
	cRand "crypto/rand"
	"math/big"
	"math/rand"
	"os"
)

type InputGetter interface {
	GetInput() (string, error)
}

type RealInputGetter struct{}

func (r *RealInputGetter) GetInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	return scanner.Text(), nil
}

func RandomInt64(len int) (int64, error) {
	n, err := cRand.Int(cRand.Reader, big.NewInt(int64(len)))
	return n.Int64(), err
}

func ScramblePassword(s string) (string, error) {
	r := []rune(s)
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})

	return string(r), nil
}
