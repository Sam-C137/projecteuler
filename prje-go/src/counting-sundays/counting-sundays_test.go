package counting_sundays

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestCountingSundays(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(171, Run())
	a.ExpectEquals(171, RunBetter())
}
