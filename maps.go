package main

import "fmt"

func main(){
	// mapの宣言1
	m1 := make(map[string]int)
	fmt.Println("m1:", m1)

	// mapの宣言2
	// 宣言と初期化を同時に行う
	m2 := map[string]int{
		"alice": 31,
		"charlie": 34,
	}
	fmt.Println("m2", m2)

	// mapに値を設定
	m := make(map[string]int)
	m["k1"] = 10
	m["k2"] = 20
	fmt.Println(m)

	// mapから値を取得
	val := m["k1"]
	fmt.Println(val)

	// mapに含まれるキーと値のペアの個数を取得
	fmt.Println("len:", len(m))

	// mapからキーと値のペアを削除
	delete(m, "k2")
	fmt.Println("deleted:", m)

	// mapから値を取得する時、存在しないキーにアクセスするとその型のゼロ値が返される
	// キーが存在して値がゼロ値なのか、キーが存在しないのかを区別するため、オプションである2つ目の返り値には
	// マップにキーが含まれているかどうかを表す真偽値が入る。
	v, ok := m["k2"]
	fmt.Println("value:", v)	// ゼロ値
	fmt.Println("is ok:", ok)	// false
}