package basic

import (
	"github.com/go-ds-algorithms/algorithms/search"
	"github.com/go-ds-algorithms/algorithms/sort"
)

// IntSlice is data structure of slice of int
type IntSlice []int

func (p IntSlice) LenFunc() int               { return len(p) }
func (p IntSlice) LessThanFunc(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) SwapFunc(i, j int)          { p[i], p[j] = p[j], p[i] }
func (p IntSlice) LinearSearch(x int) int     { return search.LinearSearchInts(p, x) }
func (p IntSlice) BinarySearch(x int) int     { return search.BinarySearchInts(p, x) }
func (p IntSlice) InsertionSort()             { sort.InsertionSortInts(0, p.LenFunc(), p) }

// Float64Slice is data structure of slice of float64
type Float64Slice []float64

func (p Float64Slice) LenFunc() int               { return len(p) }
func (p Float64Slice) LessThanFunc(i, j int) bool { return p[i] < p[j] }
func (p Float64Slice) SwapFunc(i, j int)          { p[i], p[j] = p[j], p[i] }
func (p Float64Slice) LinearSearch(x float64) int { return search.LinearSearchFloat64s(p, x) }
func (p Float64Slice) BinarySearch(x float64) int { return search.BinarySearchFloat64s(p, x) }
func (p Float64Slice) InsertionSort()             { sort.InsertionSortFloat64s(0, p.LenFunc(), p) }

// StringSlice is data structure of slice of string
type StringSlice []string

func (p StringSlice) LenFunc() int               { return len(p) }
func (p StringSlice) LessThanFunc(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) SwapFunc(i, j int)          { p[i], p[j] = p[j], p[i] }
func (p StringSlice) LinearSearch(x string) int  { return search.LinearSearchStrings(p, x) }
func (p StringSlice) BinarySearch(x string) int  { return search.BinarySearchStrings(p, x) }
func (p StringSlice) InsertionSort()             { sort.InsertionSortString(0, p.LenFunc(), p) }
