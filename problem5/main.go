// https://projecteuler.net/problem=5
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import "fmt"

func greatestCommonFactor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lowestCommonMultiple(a, b int) int {
	return (a * b) / greatestCommonFactor(a, b)
}

func main() {
	divisibleUpto := 20
	lcm := 1
	for i := 2; i <= divisibleUpto; i++ {
		lcm = lowestCommonMultiple(lcm, i)
	}

	fmt.Println(lcm)
}
