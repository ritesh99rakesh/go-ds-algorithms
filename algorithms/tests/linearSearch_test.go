package tests

import (
	"github.com/ritesh99rakesh/go-ds-algorithms/algorithms/search"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	a1 := []int{2, 4, 1, 10, 49, 3, 9, 7}
	a2 := []float64{2.2, 3.53, 10.5, 1.98, 5.1, 49.0, 7.4}
	a3 := []string{"ram", "ganesh", "krish", "laxmi", "yamuna"}
	if pos := search.LinearSearchInts(a1, 49); pos != 4 {
		t.Fail()
	}
	if pos := search.LinearSearchInts(a1, 0); pos != -1 {
		t.Fail()
	}
	if pos := search.LinearSearchFloat64s(a2, 7.4); pos != 6 {
		t.Fail()
	}
	if pos := search.LinearSearchFloat64s(a2, 8.0); pos != -1 {
		t.Fail()
	}
	if pos := search.LinearSearchStrings(a3, "laxmi"); pos != 3 {
		t.Fail()
	}
	if pos := search.LinearSearchStrings(a3, "ganga"); pos != -1 {
		t.Fail()
	}
}
