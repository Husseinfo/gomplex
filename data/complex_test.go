package data_test

import (
	"goahead/data"
	"testing"
)

func TestInvertSign(t *testing.T) {
	x := data.Complex{
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
