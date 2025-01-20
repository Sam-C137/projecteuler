package large_sum

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestLargeSum(t *testing.T) {
	a := testutils.Assert[string]{T: t}
	a.ExpectEquals("5537376230", Run())
}
