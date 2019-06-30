package search

import "errors"

/*
LinearSearch performs linear search on an array of given size
and given equality function.
Returns (pos int, err error). If element is found: returns (pos, nil),
else: returns (0, err)

It is recommended that you use LinearSearchInts, LinearSearchFloat64s
and LinearSearchStrings, as they have inbuild func and sanity check.
*/
func LinearSearch(n int, f func(int) bool) (int, error) {
	for i := 0; i < n; i++ {
		if f(i) {
			return i, nil
		}
	}
	return 0, errors.New("element not found in the array.")
}

/*
LinearSearchInts performs linear serch over int array.
Returns (pos int, err error). If element is found: returns (pos, nil),
else: returns (0, err)
*/
func LinearSearchInts(arr []int, x int) (int, error) {
	return LinearSearch(len(arr), func(i int) bool { return arr[i] == x })
}

/*
LinearSearchFloat64s performs linear serch over float64 array.
Returns (pos int, err error). If element is found: returns (pos, nil),
else: returns (0, err)
*/
func LinearSearchFloat64s(arr []float64, x float64) (int, error) {
	return LinearSearch(len(arr), func(i int) bool { return arr[i] == x })
}

/*
LinearSearchStrings performs linear serch over string array.
Returns (pos int, err error). If element is found: returns (pos, nil),
else: returns (0, err)
*/
func LinearSearchStrings(arr []string, x string) (int, error) {
	return LinearSearch(len(arr), func(i int) bool { return arr[i] == x })
}
