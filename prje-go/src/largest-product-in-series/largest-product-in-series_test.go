package largest_product_in_series

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestLargestProductInSeries_4(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(5832, Run())
}

func TestLargestProductInSeries_13(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(23514624000, Run(13))
}
