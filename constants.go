package main

import (
	"fmt"
	"math"
	"reflect"
)

// const文はvar文が書けるところならどこでも書ける
const s string = "str"

func main(){
	fmt.Println(s, reflect.TypeOf(s))

	// 数値データ型の定数は明示的にキャストするなどして型が決まるまでは型を持たない
	const n = 500000000
	fmt.Println(n, reflect.TypeOf(n))

	const d = 3e20 / n
	fmt.Println(d, reflect.TypeOf(d))

	// 型を必要とするコンテキストによって適した型が与えられる
	fmt.Println(math.Sin(n))	// math.Sinはfloat64を引数に受けとる
}

