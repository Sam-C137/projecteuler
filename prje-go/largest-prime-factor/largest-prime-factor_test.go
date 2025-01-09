package largest_prime_factor

import "testing"

func TestLargestPrimeFactor(t *testing.T) {
	if val := Run(uint64(600_851_475_143)); val != 6857 {
		t.Errorf("%d != 6857; want 6857, got %d", val, val)
	}
}
