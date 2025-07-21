package medium_level

/*
Merge Sort is a stable, comparison-based sorting algorithm that uses the divide-and-conquer strategy.
It recursively splits the input slice into halves, sorts each half, and then merges the sorted halves.

Time Complexity:
- Best Case:     O(n log n)
- Average Case:  O(n log n)
- Worst Case:    O(n log n)

Space Complexity:
- O(n) due to the use of additional slices during the merge step.

When to Use Merge Sort:
- When stable sorting is required (i.e., equal elements retain their original order).
- When working with large datasets where worst-case performance matters.
- When sorting linked lists (can be implemented with O(1) extra space).
- When sorting datasets too large to fit in memory (external sorting).

Merge Sort is not in-place and requires additional memory for merging,
making it less efficient for small datasets compared to in-place algorithms like Quick Sort.
*/
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
