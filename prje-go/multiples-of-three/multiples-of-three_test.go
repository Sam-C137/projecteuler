package multiples_of_three

import "testing"

func TestRunDefault(t *testing.T) {
	val := Run()
	if val != 23 {
		t.Errorf("%d != 23; want 23, got %d", val, val)
	}
}

func TestRunWithArgs(t *testing.T) {
	val := Run(1000)
	
	if val != 233168 {
		t.Errorf("%d != 233168; want 233168, got %d", val, val)
	}
}
