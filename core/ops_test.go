package core_test

import (
	"goahead/core"
	"goahead/data"
	"testing"
)

func TestAddComplex(t *testing.T) {
	x := data.ComplexNumber{
		Real:      2,
		Imaginary: 3,
	}

	y := data.ComplexNumber{
		Real:      -1,
		Imaginary: 2,
	}

	z := core.AddComplex(x, y)

	if z.Real != 1 {
		t.Errorf("Real sum is wrong: got %.2f instead of %d", z.Real, 1)
	}

	if z.Imaginary != 5 {
		t.Errorf("Imaginary sum is wrong: got %.2f instead of %d", z.Imaginary, 5)
	}
}

func TestSubComplex(t *testing.T) {
	x := data.ComplexNumber{
		Real:      2,
		Imaginary: 3,
	}

	y := data.ComplexNumber{
		Real:      -1,
		Imaginary: 2,
	}

	z := core.SubComplex(x, y)

	if z.Real != 3 {
		t.Errorf("Real sum is wrong: got %.2f instead of %d", z.Real, 3)
	}

	if z.Imaginary != 1 {
		t.Errorf("Imaginary sum is wrong: got %.2f instead of %d", z.Imaginary, 1)
	}
}
