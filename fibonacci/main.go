package main

import "fmt"

func fib(n int) []int {
	ret := []int{0, 1}

	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}

	return ret
}

func main() {
	n := 6

	fmt.Println(fib(n))
}
