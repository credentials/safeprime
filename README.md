# safeprime
Go binding for OpenSSL's safe prime generator.

Generating large safe prime numbers (e.g., of 2k bits) can take prohibitively long in pure Go. This library provides a Go binding for OpenSSL's much faster [`BN_generate_prime_ex`](https://www.openssl.org/docs/manmaster/crypto/BN_generate_prime.html).
