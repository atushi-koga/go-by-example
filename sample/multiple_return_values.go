package main

import "fmt"

// Goは多値の返り値を組み込みでサポートしている
func val() (int, int) {
	return 3, 7
}

func main() {
	v1, v2 := val()
	fmt.Println(v1, v2)

	// 返り値の一部だけ必要な場合はブランク識別子を使う
	v3, _ := val()
	fmt.Println(v3)
}
