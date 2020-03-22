package search

/*
LinearSearch performs linear search on an array of given size
and given equality function.
Returns (pos int). If element is found: returns pos,
else: returns -1

It is recommended that you use LinearSearchInts, LinearSearchFloat64s
and LinearSearchStrings, as they have inbuilt func and sanity check.
*/
func LinearSearch(n int, equalityFunc func(int) bool) int {
	for i := 0; i < n; i++ {
		if equalityFunc(i) {
			return i
		}
	}
	return -1
}

/*
LinearSearchInts performs linear search over int array.
Returns (pos int). If element is found: returns pos,
else: returns -1
*/
func LinearSearchInts(arr []int, x int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		panic("cannot search in array of size 0")
	}
	return LinearSearch(arrLen, func(i int) bool { return arr[i] == x })
}

/*
LinearSearchFloat64s performs linear search over float64 array.
Returns (pos int). If element is found: returns pos,
else: returns -1
*/
func LinearSearchFloat64s(arr []float64, x float64) int {
	arrLen := len(arr)
	if arrLen == 0 {
		panic("cannot search in array of size 0")
	}
	return LinearSearch(arrLen, func(i int) bool { return arr[i] == x })
}

/*
LinearSearchStrings performs linear search over string array.
Returns (pos int). If element is found: returns pos,
else: returns -1
*/
func LinearSearchStrings(arr []string, x string) int {
	arrLen := len(arr)
	if arrLen == 0 {
		panic("cannot search in array of size 0")
	}
	return LinearSearch(arrLen, func(i int) bool { return arr[i] == x })
}
