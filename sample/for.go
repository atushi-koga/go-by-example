package main

import (
	"fmt"
)

func main() {
	// 1つだけ条件を書く最も基本的な書き方
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 初期化、条件、ループ間処理を書く古典的な `for` ループ
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	// 条件を書かない `for` ループは、`break` でループから抜けるか、`return` で関数から抜けるまで繰り返し続ける。
	for {
		fmt.Println("loop")
		break
	}

	// `continue` と書くと、ループ内の次の繰り返しに進む。
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	var cards []string
	cardsSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// forループ中のrangeは、スライスもしくはマップを反復処理する
	// スライスの場合は、各繰り返しに対して2つの値が返される。1つ目はインデックスで、2つ目はそのインデックスの要素のコピー。
	// '_' で使用しないパラメータを明示
	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	fmt.Println(cards)
}
