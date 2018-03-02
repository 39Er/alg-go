package find

import "fmt"

/**
  1.将输入数组划分为n/5组，每组有5个元素，且剩下的至多有一组的元素小于5个。

  2.寻找这n/5个组中每个组的中位数，可以将每组做一次排序，然后选取每组的第三个元素。

  3.对于第2部找出的n/5个中位数递归的调用Select函数求出其中位数x.（约定偶数个中位数为其较小的中位数）

  4.按照找到的中位数x将数组划分为两个部分，求得小于或者等于x的元素有q个

  5.如果k==q则返回x,若k<q则在x的左区间找第k小的数，否则在x的有区间找第k-q大的数
*/

func WorstSelect(arr []int, key int) int {
	return find(arr, 0, len(arr)-1, key)
}

func find(arr []int, p int, q int, key int) int {
	if p == q {
		return arr[p]
	}
	if key == 0 {
		return -1
	}
	if p < q {
		midian := getMid(arr, p, q)
		index := partition(arr, p, q, midian)
		i := index - p + 1
		if i == key {
			return arr[index]
		} else if key < i {
			return find(arr, p, index-1, key)
		} else {
			return find(arr, index+1, q, key-i)
		}
	}
	return 0
}

func partition(arr []int, p int, q int, pivot int) int {
	fmt.Println("p:", p, "q:", q, "pivot:", pivot)
	pivotIdx := getIndex(arr, p, q, pivot)
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

func getIndex(arr []int, p int, q int, midian int) int {
	for i := p; i <= q; i++ {
		if arr[i] == midian {
			return i
		}
	}
	return -1
}

func getMid(arr []int, p int, q int) int {
	if p == q {
		return arr[p]
	}
	n := q - p + 1
	var m int
	if n%5 == 0 {
		m = n / 5
	} else {
		m = n/5 + 1
	}
	midArray := make([]int, m)
	index := 0
	for index = p; index <= q-5; index += 5 {
		insert_sort(arr, index, 4)
		num := index - p
		midArray[num/5] = arr[index+2]
		fmt.Println("midArray", midArray)
	}
	remain_num := q - index + 1
	if remain_num > 0 {
		insert_sort(arr, index, remain_num-1)
		mid := midPostion(remain_num)
		num := index - p
		midArray[num/5] = arr[index+mid]
		fmt.Println("midArray", midArray)
	}
	insert_sort(midArray, 0, m-1)
	return midArray[midPostion(m)]
}

func insert_sort(arr []int, p int, step int) {
	fmt.Println("p", p, "step", step)
	fmt.Println("before:", arr)
	for j := p; j <= p+step; j++ {
		e := arr[j]
		i := j - 1
		for i >= p && arr[i] > e {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = e
	}
	fmt.Println("after:", arr)
}

func midPostion(length int) int {
	if length%2 == 0 {
		return length/2 - 1
	} else {
		return length / 2
	}
}
