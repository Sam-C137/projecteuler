package _0_001st_prime

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func Test10_001stPrime(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(104743, Run(10_001))
}
