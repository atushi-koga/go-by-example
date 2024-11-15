package main

import "fmt"

// rangeは様々なデータ構造の要素の値を繰り返し取得するのに役立つ

func main() {
	// スライスから要素を取得（配列の場合もスライスの時と同じ使い方をする）
	// 返り値の1つ目はインデックス、2つ目は値が入る
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// mapからキーと値のペアを取得
	kvs := map[string]int{"k1": 10, "k2": 20, "k3": 30}
	for k, v := range kvs {
		fmt.Printf("%s => %d\n", k, v)
	}

	// 文字列からUnicodeコードポイントを取得
	for i, c := range "go" {
		fmt.Println("index:", i, "value:", c)
	}
}
