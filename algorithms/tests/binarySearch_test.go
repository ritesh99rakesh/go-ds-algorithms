package search

import (
	"github.com/go-ds-algorithms/algorithms/search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 7, 9, 10, 49}
	a2 := []float64{1.98, 2.2, 3.53, 5.1, 7.4, 10.5, 49.0}
	a3 := []string{"ganesh", "krish", "laxmi", "ram", "yamuna"}
	if pos := search.BinarySearchInts(a1, 49); pos != 7 {
		t.Fail()
	}
	if pos := search.BinarySearchInts(a1, 0); pos != -1 {
		t.Fail()
	}
	if pos := search.BinarySearchFloat64s(a2, 7.4); pos != 4 {
		t.Fail()
	}
	if pos := search.BinarySearchFloat64s(a2, 8.0); pos != -1 {
		t.Fail()
	}
	if pos := search.BinarySearchStrings(a3, "laxmi"); pos != 2 {
		t.Fail()
	}
	if pos := search.BinarySearchStrings(a3, "ganga"); pos != -1 {
		t.Fail()
	}
}
