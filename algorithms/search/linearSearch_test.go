package search

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	a1 := []int{2, 4, 1, 10, 49, 3, 9, 7}
	a2 := []float64{2.2, 3.53, 10.5, 1.98, 5.1, 49.0, 7.4}
	a3 := []string{"ram", "ganesh", "krish", "laxmi", "yamuna"}
	if pos, err := LinearSearchInts(a1, 49); pos != 4 || err != nil {
		t.Fail()
	}
	if pos, err := LinearSearchInts(a1, 0); pos != 0 || err == nil {
		t.Fail()
	}
	if pos, err := LinearSearchFloat64s(a2, 7.4); pos != 6 || err != nil {
		t.Fail()
	}
	if pos, err := LinearSearchFloat64s(a2, 8.0); pos != 0 || err == nil {
		t.Fail()
	}
	if pos, err := LinearSearchStrings(a3, "laxmi"); pos != 3 || err != nil {
		t.Fail()
	}
	if pos, err := LinearSearchStrings(a3, "ganga"); pos != 0 || err == nil {
		t.Fail()
	}
}
