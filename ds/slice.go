package ds

import (
	"github.com/ritesh99rakesh/go-ds-algorithms/algorithms/search"
	"github.com/ritesh99rakesh/go-ds-algorithms/algorithms/sort"
)

// IntSlice is data structure of slice of int
type IntSlice []int

// Len ...
func (p IntSlice) Len() int { return len(p) }

// Swap ...
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// LinearSearch ...
func (p IntSlice) LinearSearch(x int) int { return search.LinearSearchInts(p, x) }

// BinarySearch ...
func (p IntSlice) BinarySearch(x int) int { return search.BinarySearchInts(p, x) }

// InsertionSort ...
func (p IntSlice) InsertionSort() { sort.InsertionSortInts(0, p.Len(), p) }

// Max ...
func (p IntSlice) Max() (int, int) { return search.MaxInts(p) }

// Float64Slice is data structure of slice of float64
type Float64Slice []float64

// Len ...
func (p Float64Slice) Len() int { return len(p) }

// Swap ...
func (p Float64Slice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// LinearSearch ...
func (p Float64Slice) LinearSearch(x float64) int { return search.LinearSearchFloat64s(p, x) }

// BinarySearch ...
func (p Float64Slice) BinarySearch(x float64) int { return search.BinarySearchFloat64s(p, x) }

// InsertionSort ...
func (p Float64Slice) InsertionSort() { sort.InsertionSortFloat64s(0, p.Len(), p) }

// Max ...
func (p Float64Slice) Max() (float64, int) { return search.MaxFloat64s(p) }

// StringSlice is data structure of slice of string
type StringSlice []string

// Len ...
func (p StringSlice) Len() int { return len(p) }

// Swap ...
func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// LinearSearch ...
func (p StringSlice) LinearSearch(x string) int { return search.LinearSearchStrings(p, x) }

// BinarySearch ...
func (p StringSlice) BinarySearch(x string) int { return search.BinarySearchStrings(p, x) }

// InsertionSort ...
func (p StringSlice) InsertionSort() { sort.InsertionSortString(0, p.Len(), p) }

// Max ...
func (p StringSlice) Max() (string, int) { return search.MaxStrings(p) }
