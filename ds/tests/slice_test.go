package tests

import (
	"github.com/go-ds-algorithms/ds"
	"reflect"
	"testing"
)

func TestLenFunc(t *testing.T) {
	a1 := ds.IntSlice{2, 4, 1, 10, 49, 3, 9, 7}
	a1Sorted := ds.IntSlice{1, 2, 3, 4, 7, 9, 10, 49}
	a2 := ds.Float64Slice{2.2, 3.53, 10.5, 1.98, 5.1, 49.0, 7.4}
	a2Sorted := ds.Float64Slice{1.98, 2.2, 3.53, 5.1, 7.4, 10.5, 49.0}
	a3 := ds.StringSlice{"ram", "ganesh", "krish", "laxmi", "yamuna"}
	a3Sorted := ds.StringSlice{"ganesh", "krish", "laxmi", "ram", "yamuna"}
	if a1.Len() != len(a1) {
		t.Fail()
	}
	if a2.Len() != len(a2) {
		t.Fail()
	}
	if a3.Len() != len(a3) {
		t.Fail()
	}
	if a1.LinearSearch(2) != 0 {
		t.Fail()
	}
	if a2.LinearSearch(2.2) != 0 {
		t.Fail()
	}
	if a3.LinearSearch("ram") != 0 {
		t.Fail()
	}
	if a1Sorted.BinarySearch(1) != 0 {
		t.Fail()
	}
	if a2Sorted.BinarySearch(1.98) != 0 {
		t.Fail()
	}
	if a3Sorted.BinarySearch("ganesh") != 0 {
		t.Fail()
	}
	if a1.Swap(0, 1); !reflect.DeepEqual(a1, ds.IntSlice{4, 2, 1, 10, 49, 3, 9, 7}) {
		t.Fail()
	}
	if a2.Swap(0, 1); !reflect.DeepEqual(a2, ds.Float64Slice{3.53, 2.2, 10.5, 1.98, 5.1, 49.0, 7.4}) {
		t.Fail()
	}
	if a3.Swap(0, 1); !reflect.DeepEqual(a3, ds.StringSlice{"ganesh", "ram", "krish", "laxmi", "yamuna"}) {
		t.Fail()
	}
	if a1.InsertionSort(); !reflect.DeepEqual(a1, a1Sorted) {
		t.Fail()
	}
	if a2.InsertionSort(); !reflect.DeepEqual(a2, a2Sorted) {
		t.Fail()
	}
	if a3.InsertionSort(); !reflect.DeepEqual(a3, a3Sorted) {
		t.Fail()
	}
}
