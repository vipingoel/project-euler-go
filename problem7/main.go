// https://projecteuler.net/problem=7
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

// What is the 10 001st prime number?

package main

import (
	"fmt"
	"math"
)

func isPrime(n int64) bool {
	for i := int64(3); i <= int64(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	// Init
	nth := 6
	prime := int64(13)

	for i := prime; nth <= 10001; i += 2 {
		if isPrime(i) {
			prime = i
			nth++
		}
	}

	fmt.Println(prime)
}
