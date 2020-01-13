package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	// Goはメソッドを呼ぶとき、値とポインタの変換を自動で行う
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rptr := &r
	fmt.Println("area:", rptr.area())
	fmt.Println("perim:", rptr.perim())
}
