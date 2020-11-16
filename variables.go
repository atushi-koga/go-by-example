package main

import "fmt"

func main(){
	// 変数の宣言と初期化
	// 変数を初期化すれば、Goは変数の型を推測する
	var a = "initial"
	fmt.Println(a)	// initial

	// 複数の変数を一度に宣言
	var b, c = 1, 2
	fmt.Println(b, c)	// 1 2

	// 明示的に初期化していない変数の値はその型のゼロ値
	var d1 bool
	var d2 string
	var d3 int
	var d4 float64
	fmt.Printf("%v, %v, %v, %v", d1, d2, d3, d4) // false, "", 0, 0

	// 宣言と初期化を行う省略記法
	f := "short"
	fmt.Println(f)	// short

	// 上の省略記法は以下と同じ。これは冗長になるので、型を省略するのが一般的。
	var g string = "long"
	fmt.Println(g)	// long
}