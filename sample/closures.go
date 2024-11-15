package main

import "fmt"

// 関数bodyの中で定義した無名関数を返す
// 返される関数に変数iを閉じ込めており、クロージャを作っている
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// この関数値は独自のiを持っていて、コールするたびにそのiが更新される
	seqFunc := intSeq()
	fmt.Println(seqFunc()) // 1
	fmt.Println(seqFunc()) // 2

	// intSeq関数内の変数iは、呼び出されるたびに生成されるローカル変数なので、
	// seqFuncが参照するiとanotherSeqFuncが参照するiは別物であり、
	// intSeq関数を呼び出すたびに独立したシーケンス関数を生成できる
	anotherSeqFunc := intSeq()
	fmt.Println(anotherSeqFunc()) // 1
	fmt.Println(anotherSeqFunc()) // 2

	fmt.Println(seqFunc()) // 3
}
