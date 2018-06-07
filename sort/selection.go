package sort

import "fmt"

func SelectionSort(arr []int) []int {
	var minValue int
	var minIndex int
	for i := 0; i < len(arr); i++ {
		minIndex = i
		minValue = arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < minValue {
				minIndex = j
				minValue = arr[j]
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		fmt.Println(arr)
	}
	return arr
}
