package amicable_numbers

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestAmicableNumbers(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(31626, Run(10_000))
}
