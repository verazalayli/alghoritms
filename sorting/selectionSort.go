package sorting

func SelectionSort(arr []int) []int {
	sorted := make([]int, len(arr))
	length := len(arr)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
				sorted = append(sorted, arr[minIndex])
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}
