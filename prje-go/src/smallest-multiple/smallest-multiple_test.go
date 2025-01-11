package smallest_multiple

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestSmallestMultiple1_10(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(2520, Run())
}

func TestSmallestMultiple1_20(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(232792560, Run(20))
}
