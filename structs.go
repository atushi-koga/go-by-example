package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	// 構造体の初期化
	// 明記しないフィールドの値はゼロ値になる
	fmt.Println("init1:", person{name: "Bob", age: 10})
	fmt.Println("init2:", person{})
	fmt.Println("init3:", person{name: "Bob"})
	fmt.Println("init4:", person{age: 10})

	// 構造体フィールドへのアクセス
	p := person{name: "Alice", age: 20}
	fmt.Println("field:", p.name)

	// 構造体のフィールドを変更する
	p.age = 30
	fmt.Println("mutable:", p.age)

	// &演算子で構造体のポインタを作れる
	pptr := &person{name: "Mike", age: 30}

	// 構造体のポインタにドット演算子を使ってフィールドへアクセスできる
	fmt.Println("pointer access1:", pptr.age)
	pptr.name = "Michael"
	fmt.Println("pointer access2:", pptr.name)
}
