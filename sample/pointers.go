package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	// ポインタレシーバの提供1
	// アドレス演算子&: 右辺の変数のアドレスを得る
	// アドレス演算子*: 右辺の変数に格納されたアドレスを解決する
	p1 := Point{1, 2}
	pptr := &p1
	pptr.ScaleBy(2)
	fmt.Println("p1_1", p1)    // {2, 4}
	fmt.Println("p1_2", *pptr) // アドレスの解決

	// ポインタレシーバの提供2
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println("p2", *r) // {2, 4}

	// ポインタレシーバの提供3
	p3 := Point{1, 2}
	(&p3).ScaleBy(2)
	fmt.Println("p3", p3) // {2, 4}

	// ポインタレシーバの提供4
	// *Pointレシーバを要求する場合でも、Point型の変数を渡せばコンパイラが暗黙的にポインタ型に変換する（アドレス化可能なため）
	p4 := Point{1, 2}
	p4.ScaleBy(2)
	(&p4).ScaleBy(2)      // 上記と同じ
	fmt.Println("p4", p4) // {4, 8}

	// 逆に、Point型を要求する場合でも、*Pointレシーバでメソッドを呼び出すことができる （アドレスから値を得られるため）
	// つまり、レシーバが指している値をロードする
	// コンパイラは暗黙的な*演算子を挿入する
	p5 := Point{1, 2}
	p6ptr := &Point{3, 4}
	fmt.Println("p5_1", p6ptr.Distance(p5))
	fmt.Println("p5_2", (*p6ptr).Distance(p5)) // 上記と同じ
}
