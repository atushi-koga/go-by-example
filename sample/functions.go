package main

import "fmt"

// 関数宣言
// 引数、返り値に型を指定する
func plus(a int, b int) int {
	return a + b
}

// 同じ型の引数を続けて複数受け取る時、それらの型名はそのうち最後のものの後にだけ記述すれば良い
func plusPlus(a, b, c int) int {
	return a + b + c
}

// 返り値がない場合は、返り値の型宣言は書かない。(PHPで言うvoidは存在しない)
func printFmt(v int) {
	fmt.Println(v)
}

func main() {
	res := plus(1, 2)
	printFmt(res)

	res2 := plusPlus(1, 2, 3)
	printFmt(res2)
}
