// https://projecteuler.net/problem=2
// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

package main

import (
	"fmt"
)

func main() {
	res := int64(0)
	prev := int64(1)
	for next := int64(2); next < int64(4000000); prev, next = next, prev+next {
		if next%2 == 0 {
			fmt.Println(prev, ",", next)
			res += next
		}
	}

	fmt.Println(res)
}
