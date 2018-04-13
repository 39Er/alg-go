package find

func BinarySearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if arr[mid] > key {
			high = mid - 1
		} else if arr[mid] < key {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
