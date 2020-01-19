package main

import (
	"fmt"
	"math"
)

// インターフェースとは、メソッドのシグネチャの集まりに名前を付けたものである
// これは幾何学図形のインターフェース
type geometry interface {
	area() float64
	perimeter() float64
}

// rectangular型とcircle型でgeometryインターフェースを実装する
// Goのインターフェースを実装するには、インターフェースに含まれるメソッドを全て実装すれば良い
type rectangular struct{
	width, height float64
}
type circle struct{
	radius float64
}

// rectangular型でgeometryインターフェースを実装する
func (r rectangular) area() float64 {
	return r.width * r.height
}
func (r rectangular) perimeter() float64 {
	return 2 * r.width + 2 * r.height
}

// circle型でgeometryインターフェースを実装する
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 変数がインターフェース型を持つなら、そのインターフェースのメソッドを呼ぶことができる
// geometry型の変数であればこのメソッドを使える
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rectangular{width:3, height:4}
	c := circle{radius:5}
	measure(r)
	measure(c)
}
