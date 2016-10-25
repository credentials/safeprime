// Package safeprime is a small wrapper around openssl's BN_generate_prime_ex for generating safe primes.
package safeprime

/*
#cgo pkg-config: libssl
#include <openssl/bn.h>
#include <openssl/rand.h>

static char* openssl_generate_safeprime(int bitsize) {
	BIGNUM* bignum;
	bignum = BN_new();
	if (bignum == NULL)
		return NULL;

	//NOTE: we do not initialize openssl's PRNG, as it is not necessary on machines that have /dev/urandom.
	// Might want to do this for Windows machines.  See https://www.openssl.org/docs/manmaster/crypto/RAND_add.html
	if (BN_generate_prime_ex(bignum, bitsize, 1, NULL, NULL, NULL) != 1)
		return NULL;

	// Return a char array containing a hexadecimal representation of the bignum
	char* hex = BN_bn2hex(bignum);
	BN_free(bignum);
	if (hex == NULL)
		return NULL;
	return hex;
}
*/
import "C"
import (
	"errors"
	"math/big"
	"unsafe"
)

// Generate uses openssl's BN_generate_prime_ex to generate a new safe prime of the given size.
func Generate(size int) (*big.Int, error) {
	// Generate the prime
	bignum := C.openssl_generate_safeprime(C.int(size))
	defer C.free(unsafe.Pointer(bignum))
	if bignum == nil {
		return nil, errors.New("openssl failed to generate a safe prime")
	}

	// Convert the C string to a big.Int
	x := new(big.Int)
	x.SetString(C.GoString(bignum), 16)

	return x, nil
}
