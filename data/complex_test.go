package data_test

import (
	"goahead/data"
	"testing"
)

func TestInvertSign(t *testing.T) {
	x := data.ComplexNumber{
		Real:      2,
		Imaginary: 3,
	}

	inverted := x.ToInvertedSign()

	if x.Real+inverted.Real != 0 {
		t.Errorf("Wrong real number")
	}

	if x.Imaginary+inverted.Imaginary != 0 {
		t.Errorf("Wrong real imaginary")
	}
}


func TestComplexNumber_GetType(t *testing.T) {
	x := data.ComplexNumber{
		Real:      2,
		Imaginary: 0,
	}

	if x.GetType() != data.Real{
		t.Errorf("Wrong Type: expected %d, got %d", data.Real, x.GetType())
	}

	y := data.ComplexNumber{
		Real:      0,
		Imaginary: -1,
	}

	if y.GetType() != data.Imaginary{
		t.Errorf("Wrong Type: expected %d, got %d", data.Imaginary, y.GetType())
	}

	z := data.ComplexNumber{
		Real:      2,
		Imaginary: -1,
	}

	if z.GetType() != data.Complex{
		t.Errorf("Wrong Type: expected %d, got %d", data.Complex, z.GetType())
	}
}
