package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	if _, val := Run(4_000_000); val != 4613732 {
		t.Errorf("%d != 4613732; want 4613732, got %d", val, val)
	}
}
