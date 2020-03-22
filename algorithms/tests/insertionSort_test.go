package search

import (
	"github.com/go-ds-algorithms/algorithms/sort"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a1 := []int{2, 4, 1, 10, 49, 3, 9, 7}
	a1Sorted := []int{1, 2, 3, 4, 7, 9, 10, 49}
	if sort.InsertionSortInts(0, len(a1), a1); !reflect.DeepEqual(a1, a1Sorted) {
		t.Fail()
	}
	a1 = []int{2, 4, 1, 10, 49, 3, 9, 7}
	a1Sorted2 := []int{2, 1, 3, 4, 9, 10, 49, 7}
	if sort.InsertionSortInts(1, len(a1)-1, a1); !reflect.DeepEqual(a1, a1Sorted2) {
		t.Fail()
	}
	a2 := []float64{2.2, 3.53, 10.5, 1.98, 5.1, 49.0, 7.4}
	a2Sorted := []float64{1.98, 2.2, 3.53, 5.1, 7.4, 10.5, 49.0}
	if sort.InsertionSortFloat64s(0, len(a2), a2); !reflect.DeepEqual(a2, a2Sorted) {
		t.Fail()
	}
	a2 = []float64{2.2, 3.53, 10.5, 1.98, 5.1, 49.0, 7.4}
	a2Sorted2 := []float64{2.2, 1.98, 3.53, 5.1, 10.5, 49.0, 7.4}
	if sort.InsertionSortFloat64s(1, len(a2)-1, a2); !reflect.DeepEqual(a2, a2Sorted2) {
		t.Fail()
	}
	a3 := []string{"ram", "ganesh", "krish", "laxmi", "yamuna"}
	a3Sorted := []string{"ganesh", "krish", "laxmi", "ram", "yamuna"}
	if sort.InsertionSortString(0, len(a3), a3); !reflect.DeepEqual(a3, a3Sorted) {
		t.Fail()
	}
	a3 = []string{"ram", "ganesh", "krish", "laxmi", "yamuna"}
	a3Sorted2 := []string{"ram", "ganesh", "krish", "laxmi", "yamuna"}
	if sort.InsertionSortString(1, len(a3)-1, a3); !reflect.DeepEqual(a3, a3Sorted2) {
		t.Fail()
	}
}
