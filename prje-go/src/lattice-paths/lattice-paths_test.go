package lattice_paths

import (
	testutils "prje-go/helpers/test-utils"
	"testing"
)

func TestLatticePaths(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(6, Run(2))
}

func TestLatticePaths20(t *testing.T) {
	a := testutils.Assert[int]{T: t}
	a.ExpectEquals(137846528820, Run(20))
}
