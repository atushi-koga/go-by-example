package main

import (
	"fmt"
	"sort"
)

// スライスと同様にマップは比較できないので、ループを書く必要がある
func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for xk, xv := range x {
		if yv, ok := y[xk]; !ok || xv != yv {
			return false
		}
	}
	return true
}

func main() {
	// mapの宣言1: make関数で生成
	m1 := make(map[string]int)
	fmt.Println("m1:", m1)

	// mapの宣言2: マップリテラルで宣言と初期化を同時に行う
	// キーの型は == を使って比較可能でなければならない
	// 全てのキーは同じ型、全ての値は同じ型でなければならない
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println("m2:", m2)

	// mapにキーと値の組を追加
	ages := make(map[string]int)
	ages["alice"] = 10
	ages["bob"] = 20
	fmt.Println(ages)

	// mapから値を取得
	fmt.Println("alice age:", ages["alice"])

	// mapに含まれるキーと値のペアの個数を取得
	fmt.Println("ages length:", len(ages))

	// mapからキーと値のペアを削除
	delete(ages, "bob")
	fmt.Println("deleted:", ages)

	// mapから値を取得する時、存在しないキーにアクセスするとその型のゼロ値が返される
	// キーが存在して値がゼロ値なのか、キーが存在しないのかを区別するため、オプションである2つ目の返り値には
	// マップにキーが含まれているかどうかを表す真偽値が入る。
	v, ok := ages["bob"]
	fmt.Println("value:", v)  // ゼロ値
	fmt.Println("is ok:", ok) // false

	var nilMap map[string]int
	fmt.Println("nilMap equal nil:", nilMap == nil)
	fmt.Println("nilMap is 0 length:", len(nilMap) == 0)
	//nilMap["error"] = 500	// nilマップへ保存しようとするとパニックになる。マップに保存する前にマップを割り当てなければならない
	fmt.Println(nilMap)

	notNilMap := make(map[string]int)
	fmt.Println("notNilMap equal nil:", notNilMap == nil)
	fmt.Println("notNilMap is 0 length:", len(notNilMap) == 0)
	notNilMap["ok"] = 200 // マップが割り当てられているのでパニックにはならない
	fmt.Println(notNilMap)

	// mapのソート
	// マップの繰り返し順序は定義されておらず、実行ごとに変化する。
	// キーと値の組を決まった順序で列挙するためには、明示的にキーをソートしなければならない。
	ages2 := map[string]int{
		"cathy": 10,
		"alice": 20,
		"bob":   30,
	}
	var names []string
	for name := range ages2 {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages2[name])
	}

	// mapの比較
	x := map[string]int{"japan": 25, "america": 20, "china": 22}
	y := map[string]int{"japan": 25, "america": 20, "china": 22}
	z := map[string]int{"japan": 25, "america": 20}
	fmt.Println(mapEqual(x, y))
	fmt.Println(mapEqual(y, z))
}
