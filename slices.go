package main

import "fmt"

func main(){
	// スライスの宣言1
	var s1 []string
	fmt.Println("declare1:", s1)

	// スライスの宣言2
	// 長さ0でないスライスを作るにはmake関数を使う
	s2 := make([]string, 3)
	fmt.Println("declare2:", s2)

	// スライスの宣言3
	// 宣言と初期化をまとめて行う
	s3 := []string{"one", "two", "three"}
	fmt.Println("declare3:", s3)

	// 値の読み書きは配列と同様にできる
	s := make([]string, 3)
	fmt.Println("empty:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[2])
	fmt.Println("length:", len(s))

	// 配列が持っている基本操作に加えて、スライスはより豊富な操作が可能
	// 要素の追加
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	// スライスのコピー
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// スライスのスライス
	l1 := s[2:5]		// s[2],s[3],s[4]から成るスライス
	fmt.Println("slice2~4:", l1)

	// スライスのインデックス指定の省略
	l2 := s[:5]		// s[0]からs[4]で成るスライス
	fmt.Println("slice0~4:", l2)

	l3 := s[2:]    // s[2]からs[len(s)]で成るスライス
	fmt.Println("slice2~len(s):", l3)

	// スライスを組み合わせて多次元のデータ構造を作れる
	// 内側の配列の長さは要素によって異なっていても構わない。（配列の場合は、同じでなければならない）
	sd := make([][]int, 3)
	fmt.Println("first dimension:", len(sd))	// 1次元目のスライス長は3
	fmt.Println("now second dimension:", len(sd[0]))	// まだ2次元目のスライス長は0
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		// 2次元目のスライス長を設定し値を入れる
		sd[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			sd[i][j] = j
		}
	}
	fmt.Println("sd:", sd)
}
