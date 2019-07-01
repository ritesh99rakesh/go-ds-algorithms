package search

/*
LinearSearch performs linear search on an array of given size
and given equality function.
Returns (pos int). If element is found: returns pos,
else: returns -1

It is recommended that you use LinearSearchInts, LinearSearchFloat64s
and LinearSearchStrings, as they have inbuild func and sanity check.
*/
func LinearSearch(n int, f func(int) bool) int {
	for i := 0; i < n; i++ {
		if f(i) {
			return i
		}
	}
	return -1
}

/*
LinearSearchInts performs linear serch over int array.
Returns (pos int). If element is found: returns pos,
else: returns -1
*/
func LinearSearchInts(arr []int, x int) int {
	return LinearSearch(len(arr), func(i int) bool { return arr[i] == x })
}

/*
LinearSearchFloat64s performs linear serch over float64 array.
Returns (pos int). If element is found: returns pos,
else: returns -1
*/
func LinearSearchFloat64s(arr []float64, x float64) int {
	return LinearSearch(len(arr), func(i int) bool { return arr[i] == x })
}

/*
LinearSearchStrings performs linear serch over string array.
Returns (pos int). If element is found: returns pos,
else: returns -1
*/
func LinearSearchStrings(arr []string, x string) int {
	return LinearSearch(len(arr), func(i int) bool { return arr[i] == x })
}
