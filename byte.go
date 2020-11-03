package main

import "fmt"

func main() {
	// 文字列をバイトのスライスに変換
	greetingByte2 := []byte("HI")
	fmt.Println(greetingByte2)  // [72 73]

	// バイトスライスを文字列に戻す
	fmt.Println(string(greetingByte2))	// HI

	// バイトスライスをアスキーコードで生成
	greetingByte := []byte{72, 73}
	fmt.Println(greetingByte)	// [72 73]

	// バイトスライスを文字列に変換
	fmt.Println(string(greetingByte))	// HI
}
