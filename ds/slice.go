package ds

import (
	"github.com/go-ds-algorithms/algorithms/search"
	"github.com/go-ds-algorithms/algorithms/sort"
)

// IntSlice is data structure of slice of int
type IntSlice []int

func (p IntSlice) Len() int               { return len(p) }
func (p IntSlice) Swap(i, j int)          { p[i], p[j] = p[j], p[i] }
func (p IntSlice) LinearSearch(x int) int { return search.LinearSearchInts(p, x) }
func (p IntSlice) BinarySearch(x int) int { return search.BinarySearchInts(p, x) }
func (p IntSlice) InsertionSort()         { sort.InsertionSortInts(0, p.Len(), p) }
func (p IntSlice) Max() int               { return search.MaxInts(p) }

// Float64Slice is data structure of slice of float64
type Float64Slice []float64

func (p Float64Slice) Len() int                   { return len(p) }
func (p Float64Slice) Swap(i, j int)              { p[i], p[j] = p[j], p[i] }
func (p Float64Slice) LinearSearch(x float64) int { return search.LinearSearchFloat64s(p, x) }
func (p Float64Slice) BinarySearch(x float64) int { return search.BinarySearchFloat64s(p, x) }
func (p Float64Slice) InsertionSort()             { sort.InsertionSortFloat64s(0, p.Len(), p) }
func (p Float64Slice) Max() int                   { return search.MaxFloat64s(p) }

// StringSlice is data structure of slice of string
type StringSlice []string

func (p StringSlice) Len() int                  { return len(p) }
func (p StringSlice) Swap(i, j int)             { p[i], p[j] = p[j], p[i] }
func (p StringSlice) LinearSearch(x string) int { return search.LinearSearchStrings(p, x) }
func (p StringSlice) BinarySearch(x string) int { return search.BinarySearchStrings(p, x) }
func (p StringSlice) InsertionSort()            { sort.InsertionSortString(0, p.Len(), p) }
func (p StringSlice) Max() int                  { return search.MaxStrings(p) }
