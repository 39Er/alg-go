package sort

//构建大顶堆，将堆顶与最后一个元素交换
func HeapSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		buildMaxHeap(arr, length-1-i)
		arr[0], arr[length-1-i] = arr[length-1-i], arr[0] //将堆顶的最大值替换到末尾
	}
}

//构建大顶堆
func buildMaxHeap(arr []int, lastIndex int) {
	for i := (lastIndex - 1) / 2; i >= 0; i-- {
		k := i
		for k*2+1 <= lastIndex { //循环比较父子节点的大小，直到最后一个节点
			biggerIndex := k*2 + 1
			if biggerIndex < lastIndex { //存在右节点
				if arr[biggerIndex] < arr[biggerIndex+1] {
					biggerIndex++
				}
			}
			if arr[k] < arr[biggerIndex] { //子节点大于父节点
				arr[k], arr[biggerIndex] = arr[biggerIndex], arr[k] //	替换父子节点
				k = biggerIndex                                     //将父节点设为替换前的子节点，继续向下校验子节点与父节点的大小
			} else { //若父节点大于子节点
				break
			}
		}
	}
}
