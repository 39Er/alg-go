package find

import (
	"fmt"
	"math/rand"
)

func RandSelect1(arr []int, key int) int {
	_, value := quick(arr, 0, len(arr)-1, key)
	return value
}

func quick(arr []int, p int, q int, key int) (int, int) {
	for p <= q {
		index := arrange(arr, p, q) + 1
		fmt.Println(arr)
		fmt.Println("index", index)
		fmt.Println(key)
		if index == key {
			fmt.Println("aaaaaaaaaaaaaaaa")
			return key, arr[index-1]
		} else if index < key {
			fmt.Println("bbbbbbbbbbb")
			arr = arr[index:]
			fmt.Println("arr", arr)
			key = key - index
			fmt.Println("key", key)
			return quick(arr, 0, len(arr)-1, key)
		} else {
			fmt.Println("ccccccccccccccc")
			arr = arr[:index-1]
			fmt.Println("arr", arr)
			//key = index - key
			fmt.Println("key", key)
			return quick(arr, p, len(arr)-1, key)
		}
	}
	return 0, 0
}

func arrange(arr []int, p int, q int) int {
	temp := arr[p]
	for p < q {
		for p < q && arr[q] >= temp {
			q--
		}
		arr[p] = arr[q]
		for p < q && arr[p] < temp {
			p++
		}
		arr[q] = arr[p]
	}
	arr[p] = temp
	return p
}

func RandSelect2(arr []int, key int) int {
	return random_selection(arr, 0, len(arr)-1, key)
}

func random_selection(arr []int, p int, q int, key int) int {
	fmt.Println("p:", p, "q:", q, "key:", key)
	if p == q {
		if key != 1 {
			return 0
		} else {
			return arr[p]
		}
	}
	//if key == 0 {
	//	return -1
	//}
	if p < q {
		mid := random_partition(arr, p, q)
		i := mid - p + 1
		if i == key {
			return arr[mid]
		} else if key < i {
			return random_selection(arr, p, mid-1, key)
		} else {
			return random_selection(arr, mid+1, q, key-i)
		}
	}
	fmt.Println("end---------")
	return 0
}

func random_partition(arr []int, p int, q int) int {
	pivotIdx := p + rand.Intn(q-p+1)
	fmt.Println("pivotIdx", pivotIdx)
	pivot := arr[pivotIdx]
	fmt.Println("pivot", pivot)
	arr[q], arr[pivotIdx] = arr[pivotIdx], arr[q]
	pivotIdx = q
	i := p - 1
	for j := p; j <= q-1; j++ {
		if arr[j] <= pivot {
			i = i + 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[pivotIdx] = arr[pivotIdx], arr[i+1]
	fmt.Println("arr", arr)
	return i + 1
}
