package main

import "fmt"

//杨辉三角
/*
				1
			1		1
		1		2		1
	1		3		3		1
1		4		6		4		1
*/
func main() {
	n := 6
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, i+1)
		if i == 0 {
			arr[i][0] = 1
			continue
		}
		if i == 1 {
			arr[i][0], arr[i][1] = 1, 1
			continue
		}
		arr[i][0] = 1
		for j := 1; j < i; j++ {
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
		}
		arr[i][i] = 1
	}
	fmt.Println(arr)
}
