package main

import (
	"fmt"
	"time"
)

func main(){
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

	// switchは値だけでなく型について分岐することもできる
	// これを使ってinterfaceの型を調べることができる
	whatAmI := func(i interface{}){
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
