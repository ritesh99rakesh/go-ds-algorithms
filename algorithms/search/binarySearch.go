package search

/*
BinarySearch performs binary search on a given array and given element.
Returns (pos int). If element is found: returns pos,
else: returns -1
*/
func BinarySearch(arr []int, x int) int {
	low, high := 0, len(arr)
	for low < high {
		mid := low + (high-low)/2
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
