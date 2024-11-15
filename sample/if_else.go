package main

import (
	"fmt"
)

func main() {
	// 基本的な例
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// elseがなくても良い
	if 8%4 == 0 {
		fmt.Println("8 is divided by 4")
	}

	// 簡易ステートメント
	// 条件式の前に、簡易なステートメントを記述できる。この場合、セミコロンで区切って条件式の前に記述します。
	// 簡易ステートメントで宣言された変数は、ifのスコープ内でのみ有効（いずれの分岐からも見える）だが、if文から外れると見えなくなる。
	i := 5
	if i := 20; i < 0 {
		fmt.Println(i, "is negative")
	} else if i < 10 {
		fmt.Println(i, "is smaller than 10")
	} else {
		fmt.Println(i, "is larger than 10")
	}
	fmt.Println(i) // 5
}
