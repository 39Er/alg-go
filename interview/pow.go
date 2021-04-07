package main

import (
	"fmt"
	"math"
)

//幂运算
//重点在于：边界限制、效率优化(利用历史数据、位运算)
func main() {
	x := 5
	y := -2
	fmt.Println(pow(float64(x), y))
	fmt.Println(math.Pow(float64(x), float64(y)))
}

func pow(x float64, n int) float64 {
	//注意边界判断
	//负数、0值
	if n == 0 {
		return 1
	}
	if x == 0 && n < 0 {
		panic("error")
	}

	absN := n
	if n < 0 {
		absN = -n
	}

	//初级 O(n)
	//result := 1.0
	//for i := 1; i <= absN; i++ {
	//	result *= x
	//}
	//高级 O(logn)
	result := fib(x, absN)

	if n < 0 {
		result = 1.0 / result
	}

	return result
}

func fib(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	result := fib(x, n>>1)
	result *= result
	if n&1 == 1 {
		result *= x
	}
	return result
}
