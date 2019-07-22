package tests

import (
	"github.com/go-ds-algorithms/ds/basic"
	"reflect"
	"testing"
)

func TestLenFunc(t *testing.T) {
	a1 := basic.IntSlice{2, 4, 1, 10, 49, 3, 9, 7}
	a1Sorted := basic.IntSlice{1, 2, 3, 4, 7, 9, 10, 49}
	a2 := basic.Float64Slice{2.2, 3.53, 10.5, 1.98, 5.1, 49.0, 7.4}
	a2Sorted := basic.Float64Slice{1.98, 2.2, 3.53, 5.1, 7.4, 10.5, 49.0}
	a3 := basic.StringSlice{"ram", "ganesh", "krish", "laxmi", "yamuna"}
	a3Sorted := basic.StringSlice{"ganesh", "krish", "laxmi", "ram", "yamuna"}
	if a1.LenFunc() != len(a1) {
		t.Fail()
	}
	if a2.LenFunc() != len(a2) {
		t.Fail()
	}
	if a3.LenFunc() != len(a3) {
		t.Fail()
	}
	if a1.LessThanFunc(0, 1) != true {
		t.Fail()
	}
	if a2.LessThanFunc(0, 1) != true {
		t.Fail()
	}
	if a3.LessThanFunc(1, 0) != true {
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
	if a1.SwapFunc(0, 1); !reflect.DeepEqual(a1, basic.IntSlice{4, 2, 1, 10, 49, 3, 9, 7}) {
		t.Fail()
	}
	if a2.SwapFunc(0, 1); !reflect.DeepEqual(a2, basic.Float64Slice{3.53, 2.2, 10.5, 1.98, 5.1, 49.0, 7.4}) {
		t.Fail()
	}
	if a3.SwapFunc(0, 1); !reflect.DeepEqual(a3, basic.StringSlice{"ganesh", "ram", "krish", "laxmi", "yamuna"}) {
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
