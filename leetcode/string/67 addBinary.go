package string

import "math/big"

func addBinary(a string, b string) string {
	ai, _ := new(big.Int).SetString(a, 2)
	bi, _ := new(big.Int).SetString(b, 2)
	ai.Add(bi, ai)
	return ai.Text(2)
}
