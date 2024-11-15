package main

import (
	"fmt"
	"time"
)

func main() {
	// 基本的な使い方
	// 定数で値を比較する。caseには式（i > 0 など）を設定できない。
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

	// カンマを使ってcaseに複数の値を設定できる
	now := time.Now().Weekday()
	switch now {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	// switchの直後に比較元を書かなければ、それは単なるif/else文
	// 定数ではないものを使ってcase式が書ける
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	// fall through はされないのが基本
	j := 15
	switch {
	case j < 10:
		fmt.Println("j is less than 10")
	case j < 20:
		fmt.Println("j is less than 20") // これだけ表示される
	case j < 30:
		fmt.Println("j is less than 30")
	default:
		fmt.Println("j is over 30")
	}

	// 推奨されないと思うが、fall through する場合は明示的に記述する
	// そうすると、次のcase文を条件判定なしで実行する
	k := 15
	switch {
	case k < 10:
		fmt.Println("k is less than 10")
	case k < 20:
		fmt.Println("k is less than 20")
		fallthrough
	case k > 100:
		fmt.Println("30 < k 100")
	default:
		fmt.Println("k is over 30")
	}

	// switchは値だけでなく型について分岐することもできる
	// これを使ってinterfaceの型を調べることができる
	whatAmI := func(i interface{}) {
		switch x := i.(type) {
		case bool:
			fmt.Println("type bool")
		case int:
			fmt.Println("type int")
		default:
			fmt.Printf("type %T\n", x)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
