package sorting

func Quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	sortedLeft := Quicksort(left)
	sortedRight := Quicksort(right)

	return append(append(sortedLeft, pivot), sortedRight...)
}
