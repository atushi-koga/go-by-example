package main

import "fmt"

func f(from string)	{
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine1")
	go f("goroutine2")

	// 無名関数に対してゴルーチンを開始できる
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}