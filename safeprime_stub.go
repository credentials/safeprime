// +build android ios windows

package safeprime

import (
	"math/big"
)

func Generate(bitsize int) (*big.Int, error) {
	panic("Not implemented on mobile platforms")
}
