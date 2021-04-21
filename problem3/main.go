// https://projecteuler.net/problem=3
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"fmt"
	"math"
)

func factors(n int64) []int64 {
	for i := int64(2); i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return append(factors(n/i), i)
		}
	}

	return []int64{n}
}

func main() {
	n := int64(600851475143)
	fs := factors(n)
	fmt.Println(fs[0])
}
