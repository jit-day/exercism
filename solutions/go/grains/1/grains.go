package grains

import (
	"errors"
	"math"
	"math/big"
)

// Square returns the square of given number
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid index")
	}
	return uint64(math.Pow(2, float64(n-1))), nil
}

// Total returns sum of following sequence:
// 1 + 2 + 4 + 8 + ... + 2^63
func Total() uint64 {
	base := big.NewInt(2)
	result := base.Exp(base, big.NewInt(64), nil)
	result = result.Sub(result, big.NewInt(1))

	return result.Uint64()
}
