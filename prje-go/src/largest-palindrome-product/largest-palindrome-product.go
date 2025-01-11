package largest_palindrome_product

import (
	"slices"
	"strconv"
)

// Run A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 x 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
func Run(target int) int {
	largest := 0

	for i := target; i > 0; i-- {
		for j := i; j > 0; j-- {
			prod := i * j

			if prod < largest {
				break
			}

			if isPalindrome(prod) {
				largest = prod
			}
		}
	}

	return largest
}

func isPalindrome(num int) bool {
	strNum := strconv.Itoa(num)
	strArr := []rune(strNum)
	slices.Reverse(strArr)

	return strNum == string(strArr)
}
