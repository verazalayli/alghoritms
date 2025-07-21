package medium_level

/*
Quick Sort is an efficient, comparison-based, in-place sorting algorithm that follows the divide-and-conquer paradigm.
It works by selecting a pivot element, partitioning the array around the pivot,
and recursively sorting the subarrays.

Time Complexity:
- Best Case:     O(n log n)        (balanced partitions)
- Average Case:  O(n log n)
- Worst Case:    O(n²)             (highly unbalanced partitions, e.g., sorted input with poor pivot choice)

Space Complexity:
- O(log n) on average due to recursion stack (in-place sorting)

When to Use Quick Sort:
- When fast average-case performance is desired
- When working with arrays (due to in-place sorting and cache efficiency)
- When memory usage needs to be minimized (compared to Merge Sort)

Quick Sort is not stable by default (equal elements may not retain original order).
Its worst-case performance can be mitigated by using randomized pivot selection or the median-of-three strategy.
*/

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
