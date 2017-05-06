// Package safeprime is a small wrapper around openssl's BN_generate_prime_ex for generating safe primes.
package safeprime

import (
	"crypto/rand"
	"math/big"
)

var bnNew func() uintptr
var bnFree func(uintptr)
var bnGenPrime func(uintptr, int, int, uintptr, uintptr, uintptr) int
var bnToHex func(uintptr) string

// Generate tries to use openssl's BN_generate_prime_ex to generate a new safe prime of the given size;
// if that fails it uses a pure Go algorithm.
func Generate(bitsize int) (*big.Int, error) {
	// num, err := GenUsingOpenssl(bitsize)
	//
	// if err != nil {
	// 	println("WARNING: using openssl failed, switching to (slower) Go algorithm")
	return GenUsingGo(bitsize)
	// }

	// return num, nil
}

// GenUsingGo generates a safe prime in pure Go.
func GenUsingGo(bitsize int) (*big.Int, error) {
	p2 := new(big.Int)

	for {
		p, err := rand.Prime(rand.Reader, bitsize)
		if err != nil {
			return nil, err
		}

		p2.Rsh(p, 1) // p2 = (p - 1)/2
		if p2.ProbablyPrime(20) {
			return p, nil
		}
	}
}
