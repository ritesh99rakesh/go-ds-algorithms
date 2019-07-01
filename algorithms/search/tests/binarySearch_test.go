package search

import (
	"github.com/go-ds-algorithms/algorithms/search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a1 := []int{2, 4, 5, 7, 20, 49, 10}
	if pos := search.BinarySearch(a1, 49); pos != 5 {
		t.Fail()
	}
	if pos := search.BinarySearch(a1, 4); pos != 1 {
		t.Fail()
	}
	if pos := search.BinarySearch(a1, 0); pos != -1 {
		t.Fail()
	}
}
