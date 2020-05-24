package search

/*
MinInts performs linear search for Minimum element on a given int array.
Returns (pos int). If more than element are Minimum returns the
first position of occurrence.
*/
func MinInts(arr []int) (int, int) {
	arrLen := len(arr)
	if arrLen == 0 {
		panic("cannot search in array of size 0")
	}
	minValue, minPos := arr[0], 0
	for currPos, currVal := range arr {
		if minValue > currVal {
			minValue, minPos = currVal, currPos
		}
	}
	return minValue, minPos
}

/*
MinFloat64s performs linear search for Minimum element on a given float64 array.
Returns (pos int). If more than element are Minimum returns the
first position of occurrence.
*/
func MinFloat64s(arr []float64) (float64, int) {
	arrLen := len(arr)
	if arrLen == 0 {
		panic("cannot search in array of size 0")
	}
	minValue, minPos := arr[0], 0
	for currPos, currVal := range arr {
		if minValue > currVal {
			minValue, minPos = currVal, currPos
		}
	}
	return minValue, minPos
}

/*
MinStrings performs linear search for Minimum element on a given string array.
Returns (pos int). If more than element are Minimum returns the
first position of occurrence.
*/
func MinStrings(arr []string) (string, int) {
	arrLen := len(arr)
	if arrLen == 0 {
		panic("cannot search in array of size 0")
	}
	minValue, minPos := arr[0], 0
	for currPos, currVal := range arr {
		if minValue > currVal {
			minValue, minPos = currVal, currPos
		}
	}
	return minValue, minPos
}
