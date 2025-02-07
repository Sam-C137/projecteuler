package factorial_digit_sum

import (
	"math/big"
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestFactorialDigitSum10(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(27, Run(big.NewInt(10)))
}

func TestFactorialDigitSum100(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(648, Run(big.NewInt(100)))
}
