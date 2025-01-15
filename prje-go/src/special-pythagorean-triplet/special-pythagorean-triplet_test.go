package special_pythagorean_triplet

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestSpecialPythagoreanTriplets(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(31875000, Run(1000))
}
