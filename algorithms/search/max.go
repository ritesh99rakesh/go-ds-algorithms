package search

/*
MaxInts performs linear search for maximum element on a given int array.
Returns (pos int). If more than element are maximum returns the
first position of occurrence.
*/
func MaxInts(arr []int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		panic("cannot search in array of size 0")
	}
	maxValue, maxPos := arr[0], 0
	for currPos, currVal := range arr {
		if maxValue < currVal {
			maxValue, maxPos = currVal, currPos
		}
	}
	return maxPos
}

/*
MaxFloat64s performs linear search for maximum element on a given float64 array.
Returns (pos int). If more than element are maximum returns the
first position of occurrence.
*/
func MaxFloat64s(arr []float64) int {
	arrLen := len(arr)
	if arrLen == 0 {
		panic("cannot search in array of size 0")
	}
	maxValue, maxPos := arr[0], 0
	for currPos, currVal := range arr {
		if maxValue < currVal {
			maxValue, maxPos = currVal, currPos
		}
	}
	return maxPos
}

/*
MaxStrings performs linear search for maximum element on a given string array.
Returns (pos int). If more than element are maximum returns the
first position of occurrence.
*/
func MaxStrings(arr []string) int {
	arrLen := len(arr)
	if arrLen == 0 {
		panic("cannot search in array of size 0")
	}
	maxValue, maxPos := arr[0], 0
	for currPos, currVal := range arr {
		if maxValue < currVal {
			maxValue, maxPos = currVal, currPos
		}
	}
	return maxPos
}
