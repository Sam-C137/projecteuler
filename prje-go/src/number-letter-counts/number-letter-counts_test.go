package number_letter_counts

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestNumberLetterCounts(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(21124, Run(1000))
}
