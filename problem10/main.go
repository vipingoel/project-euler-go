// https://projecteuler.net/problem=10
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.

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
	sum := int64(2)
	for i := int64(3); i < 2000000; i += 2 {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
