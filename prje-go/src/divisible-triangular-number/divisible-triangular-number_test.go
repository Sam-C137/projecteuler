package divisible_triangular_number

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestHighlyDivisibleTriangularNumber5(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(28, Run(5))
}

func TestHighlyDivisibleTriangularNumber500(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(76576500, Run(500))
}
