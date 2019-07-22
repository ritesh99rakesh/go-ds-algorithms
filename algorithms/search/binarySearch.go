package search

/*
BinarySearch performs binary search on a given array and given element.
Returns (pos int). If element is found: returns pos,
else: returns -1

It is recommended that you use BinarySearchInts, BinarySearchFloat64s
and BinarySearchStrings, as they have inbuilt func and sanity check.
*/
func BinarySearch(n int, equalityFunc func(int) bool, lessThanFunc func(int) bool) int {
	low, high := 0, n
	for low <= high {
		mid := low + (high-low)/2
		if equalityFunc(mid) {
			return mid
		} else if lessThanFunc(mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

/*
BinarySearchInts performs binary search over int array.
Returns (pos int). If element is found: returns pos,
else: returns -1
*/
func BinarySearchInts(arr []int, x int) int {
	return BinarySearch(len(arr), func(i int) bool { return arr[i] == x }, func(i int) bool { return arr[i] < x })
}

/*
BinarySearchFloat64s performs binary search over float64 array.
Returns (pos int). If element is found: returns pos,
else: returns -1
*/
func BinarySearchFloat64s(arr []float64, x float64) int {
	return BinarySearch(len(arr), func(i int) bool { return arr[i] == x }, func(i int) bool { return arr[i] < x })
}

/*
BinarySearchStrings performs binary search over string array.
Returns (pos int). If element is found: returns pos,
else: returns -1
*/
func BinarySearchStrings(arr []string, x string) int {
	return BinarySearch(len(arr), func(i int) bool { return arr[i] == x }, func(i int) bool { return arr[i] < x })
}
