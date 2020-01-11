package main

import (
	"fmt"
)

func main(){
	// 基本的な例
	if 7 % 2 == 0{
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}

	// elseがなくても良い
	if 8 % 4 == 0 {
		fmt.Println("8 is divided by 4")
	}

	// 条件の前に文を書いても良い
	// この文で宣言された変数は、いずれの分岐からも見えるが、if文から外れると見えなくなる。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	}else if num < 10 {
		fmt.Println(num, "has 1 digit")
	}else{
		fmt.Println(num, "has multiple digit")
	}
	//fmt.Println(num)		// undefined: num
}
