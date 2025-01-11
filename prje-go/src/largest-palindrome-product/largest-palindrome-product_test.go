package largest_palindrome_product

import (
	"testing"
)

func TestLargestPalindromeProduct(t *testing.T) {
	if val := Run(999); val != 906609 {
		t.Errorf("%d != 999; want 999, got %d", val, val)
	}
}
