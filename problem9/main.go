// https://projecteuler.net/problem=9
// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import "fmt"

func main() {
	sum := 1000
	for a := 1; a < sum-2; a++ {
		bPlusC := sum - a

		for b := a; b < bPlusC-1; b++ {
			c := sum - a - b
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
