package sort

func QuickSort(arr []int) {
	if len(arr) > 0 {
		sort(arr, 0, len(arr)-1)
	}
}

func sort(arr []int, left int, right int) {
	if left < right {
		mid := getMiddle(arr, left, right)
		sort(arr, left, mid)
		sort(arr, mid+1, right)
	}
}

func getMiddle(arr []int, left int, right int) int {
	temp := arr[left] //基准元素
	for left < right {
		for left < right && arr[right] >= temp {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] < temp {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = temp
	return left
}
