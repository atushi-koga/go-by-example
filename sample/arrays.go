package main

import "fmt"

func main() {
	// 配列の宣言。要素数と型を指定する
	// 配列を宣言すると要素の値は型のゼロ値になる
	var array1 [5]int
	fmt.Println(array1)

	// 配列要素に値を設定、値を取得する
	array1[4] = 100
	fmt.Println("set:", array1)
	fmt.Println("get:", array1[4])

	// 配列の長さを取得する
	fmt.Println(len(array1))

	// 配列の宣言と初期化をまとめて行う
	array2 := [5]int{0, 1, 2, 3, 4}
	fmt.Println(array2)

	// 配列は1次元だが、配列型を組み合わせて多次元のデータ構造を作れる
	var array3 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			array3[i][j] = i + j
		}
	}
	fmt.Println(array3)
}
