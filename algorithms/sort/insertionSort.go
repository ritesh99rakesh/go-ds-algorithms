package sort

/*
InsertionSort performs insertion sort on a given array and given range.
Performs an in place sort with no return

Parameters:
===========

start: Start of sorting, this position is also included in sorted array
end: End of sorting, this position is not included in sorted array
compare: Function to compare between different elements, can be used to
		sort in ascending or descending order
swap: Function to swap elements

It is recommended that you use InsertionSortInts, InsertionSortFloat64s
and InsertionSortString, as they have inbuilt func and sanity check.
*/
func InsertionSort(start, end int, compare func(int, int) bool, swap func(int, int)) {
	for i := start + 1; i < end; i++ {
		for j := i; j > start && compare(j, j-1); j-- {
			swap(j, j-1)
		}
	}
}

/*
InsertionSortInts performs sort on integer array with given range.
Performs an in place sort with no return

Parameters:
===========

start: Start of sorting, this position is also included in sorted array
end: End of sorting, this position is not included in sorted array
*/
func InsertionSortInts(start, end int, arr []int) {
	InsertionSort(start, end,
		func(i int, j int) bool { return arr[i] < arr[j] }, func(i int, j int) { arr[i], arr[j] = arr[j], arr[i] })
}

/*
InsertionSortFloat64s performs sort on float64 array with given range.
Performs an in place sort with no return

Parameters:
===========

start: Start of sorting, this position is also included in sorted array
end: End of sorting, this position is not included in sorted array
*/
func InsertionSortFloat64s(start, end int, arr []float64) {
	InsertionSort(start, end,
		func(i int, j int) bool { return arr[i] < arr[j] }, func(i int, j int) { arr[i], arr[j] = arr[j], arr[i] })
}

/*
InsertionSortString performs sort on string array with given range.
Performs an in place sort with no return

Parameters:
===========

start: Start of sorting, this position is also included in sorted array
end: End of sorting, this position is not included in sorted array
*/
func InsertionSortString(start, end int, arr []string) {
	InsertionSort(start, end,
		func(i int, j int) bool { return arr[i] < arr[j] }, func(i int, j int) { arr[i], arr[j] = arr[j], arr[i] })
}
