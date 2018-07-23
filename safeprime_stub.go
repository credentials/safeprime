// +build !cgo android ios windows

package safeprime

import (
	"math/big"
)

func Generate(bitsize int) (*big.Int, error) {
	panic("Safe prime generation is disabled")
}
