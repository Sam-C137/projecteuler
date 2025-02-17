package names_scores

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestNamesScores(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(871198282, Run())
}
