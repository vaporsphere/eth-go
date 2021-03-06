package ethutil

import (
	"math/big"
)

// Big pow
//
// Returns the power of two big integers
func BigPow(a, b int) *big.Int {
	c := new(big.Int)
	c.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), big.NewInt(0))

	return c
}

// Big
//
// Shortcut for new(big.Int).SetString(..., 0)
func Big(num string) *big.Int {
	n := new(big.Int)
	n.SetString(num, 0)

	return n
}

// BigD
//
// Shortcut for new(big.Int).SetBytes(...)
func BigD(data []byte) *big.Int {
	n := new(big.Int)
	n.SetBytes(data)

	return n
}

// Big to bytes
//
// Returns the bytes of a big integer with the size specified by **base**
// Attempts to pad the byte array with zeros.
func BigToBytes(num *big.Int, base int) []byte {
	ret := make([]byte, base/8)

	if len(num.Bytes()) > base/8 {
		return num.Bytes()
	}

	return append(ret[:len(ret)-len(num.Bytes())], num.Bytes()...)
}

// Big copy
//
// Creates a copy of the given big integer
func BigCopy(src *big.Int) *big.Int {
	return new(big.Int).Set(src)
}

// Big max
//
// Returns the maximum size big integer
func BigMax(x, y *big.Int) *big.Int {
	if x.Cmp(y) <= 0 {
		return y
	}

	return x
}

// Big min
//
// Returns the minimum size big integer
func BigMin(x, y *big.Int) *big.Int {
	if x.Cmp(y) >= 0 {
		return y
	}

	return x
}
