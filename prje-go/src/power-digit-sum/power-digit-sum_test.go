package power_digit_sum

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestPowerDigitSum(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(26, Run(15))
}

func TestPowerDigitSum1000(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(1366, Run(1000))
}
