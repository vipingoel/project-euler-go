// https://projecteuler.net/problem=4
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
)

func isPalindrome(n int64) bool {
	nStr := fmt.Sprint(n)
	nLen := len(nStr)
	for begin, end := 0, nLen-1; begin < nLen/2; begin, end = begin+1, end-1 {
		if nStr[begin] != nStr[end] {
			return false
		}
	}

	return true
}

func findMaxNumberOfNDigits(n int) int64 {
	res := int64(1)
	for i := 0; i < n; i++ {
		res *= 10
	}

	return res - 1
}

func main() {
	n := 3
	maxNumberOfNDigits := findMaxNumberOfNDigits(n)
	maxNumberOfNMinus1Digits := findMaxNumberOfNDigits(n - 1)
	maxPalindrome := int64(0)

	for i := maxNumberOfNDigits; i > maxNumberOfNMinus1Digits && i*i > maxPalindrome; i-- {
		for j := i; j > maxNumberOfNMinus1Digits; j-- {
			if i*j > maxPalindrome && isPalindrome(i*j) {
				maxPalindrome = i * j
			}
		}
	}

	fmt.Println(maxPalindrome)
}
