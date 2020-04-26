package main

import "fmt"

// type宣言は既存の型と同じ基底型を持つ新たな名前付き型を定義する
// CelsiusとFahrenheitは同じ基底型であるfloat64を持つが、同じ型ではない
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius    = -273.15
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingC     Celsius    = 0
	BoilingC      Celsius    = 100
)

func (c Celsius) toF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// Stringメソッドは、fmtパッケージによって文字列として表示される場合の表現方法を制御する
func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func main() {
	// 比較したい場合は型を合わせる
	//fmt.Println(AbsoluteZeroC == AbsoluteZeroF) // コンパイルエラー：型不一致
	fmt.Println(Fahrenheit(AbsoluteZeroC) == AbsoluteZeroF)

	// 基底型の値を使う場合は型変換で取り出す
	fmt.Println(float64(AbsoluteZeroC) == float64(AbsoluteZeroF))

	// 操作の集まりは基底型が直接使われる場合と同じ。
	// float64に対して算術演算子を適用できるように、CelsiusとFahrenheitにも適用できる
	fmt.Println(BoilingC - FreezingC) // 100
	fmt.Println(BoilingC - 0)         // 100
	fmt.Println(BoilingC - 0.0)       // 100
	//fmt.Println(BoilingC.toF() - FreezingC)	// コンパイルエラー：型不一致

	// 比較演算も同様
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	//fmt.Println(c == f)          // コンパイルエラー：型不一致
	fmt.Println(c == Celsius(f)) // true

	// 名前付き型は、その型の値に対して新たな振る舞いを定義できる
	fmt.Println(FreezingC.toF())
	fmt.Println(FreezingC) // Stringメソッドを明示的に呼び出さなくて良い
}
