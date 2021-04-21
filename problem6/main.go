// https://projecteuler.net/problem=6
// The sum of the squares of the first ten natural numbers is,

// The square of the sum of the first ten natural numbers is,

// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is .

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import "fmt"

func main() {
	n := 100
	sumOfSquares := int64(0)
	sum := 0
	for i := 1; i <= n; i++ {
		sumOfSquares += int64(i * i)
		sum += i
	}

	fmt.Println(int64(sum*sum) - sumOfSquares)
}
