package sum_square_diff

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestSumSquareDifference_10(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(2640, Run())
}

func TestSumSquareDifference_100(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(25164150, Run(100))
}
