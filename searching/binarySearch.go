package searching

/*
Binary search works only on sorted slice(array)
Bianry search take the mid of the entire slice,
then mid of the new slice(half of the entire slice) and so on,
before it will find input number
*/
func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var mid int

	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		}
		if nums[mid] < target {
			low = mid + 1
		}
	}
	return -1
}
