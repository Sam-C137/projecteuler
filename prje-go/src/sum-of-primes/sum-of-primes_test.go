package sum_of_primes

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestSumOfPrimesUnder10(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(17, Run())
}

func TestSumOfPrimes(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(142913828922, Run(2_000_000))
}
