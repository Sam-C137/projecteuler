package largest_product_in_grid

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestLargestProductInGrid(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(70600674, Run())
}
