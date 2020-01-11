package main

import "fmt"

// 可変長引数関数の宣言
func sum (nums ...int) {
	fmt.Println("args:", nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
}

func main() {
	sum(1)
	sum(1, 2)

	nums := []int{1, 2, 3}
	sum(nums...)	// スライスの展開
}
