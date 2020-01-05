package main

import "fmt"

func main(){
	// 変数の宣言と初期化
	// 変数を初期化すれば、Goは変数の型を推測する
	var a = "initial"
	fmt.Println(a)

	// 複数の変数を一度に宣言
	var b, c = 1, 2
	fmt.Println(b, c)

	// 明示的に初期化していない変数の値はその型のゼロ値
	var d bool
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// 宣言と初期化を行う省略記法
	f := "short"	// var f string = "short
	fmt.Println(f)

	// 上の省略記法は以下と同じ
	var g string = "long"
	fmt.Println(g)
}