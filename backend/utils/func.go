package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func RandString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	symbols := big.NewInt(int64(len(alphanum)))
	states := big.NewInt(0)
	states.Exp(symbols, big.NewInt(int64(n)), nil)
	r, err := rand.Int(rand.Reader, states)
	if err != nil {
		panic(err)
	}
	var bytes = make([]byte, n)
	r2 := big.NewInt(0)
	symbol := big.NewInt(0)
	for i := range bytes {
		r2.DivMod(r, symbols, symbol)
		r, r2 = r2, r
		bytes[i] = alphanum[symbol.Int64()]
	}
	return string(bytes)
}

func ParseFloatPercent(s string, bitSize int) (f float64, err error) {
	i := strings.Index(s, "%")
	if i < 0 {
		return 0, fmt.Errorf("ParseFloatPercent: percentage sign not found")
	}
	f, err = strconv.ParseFloat(s[:i], bitSize)
	if err != nil {
		return 0, err
	}
	return f / 100, nil
}
