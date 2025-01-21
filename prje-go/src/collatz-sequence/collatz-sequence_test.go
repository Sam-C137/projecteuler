package collatz_sequence

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestCollatzSequence(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(837799, Run(1_000_000))
}
