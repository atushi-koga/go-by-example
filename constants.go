package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

// const文はvar文が書けるところならどこでも書ける
const s string = "str"

const (
	noDelay time.Duration = 0               // 型指定が可能
	timeout               = 5 * time.Minute // 型指定がなければ右辺の式から推測される(この場合はtime.Duration型)
)

type Weekday int

func main() {
	fmt.Println(s, reflect.TypeOf(s)) // str string

	// 数値データ型の定数は明示的にキャストするなどして型が決まるまでは型を持たない
	const n = 500000000
	fmt.Println(n, reflect.TypeOf(n))

	const m = 3e20 / n
	fmt.Println(m, reflect.TypeOf(m))

	// 型を必要とするコンテキストによって適した型が与えられる
	fmt.Println(math.Sin(n)) // math.Sinはfloat64を引数に受けとる

	fmt.Printf("%T %[1]v\n", noDelay) // time.Duration 0s
	fmt.Printf("%T %[1]v\n", timeout) // time.Duration 5m0

	// 一連の定数をグループとして宣言する場合には、グループの最初の定数を除いてすべての定数に対して
	// 右辺の式を省略することができ、前の式とその型を再び使う事を意味する
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"

	// const宣言では定数生成器iotaが使える
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday) // 0 1 2 3 4 5 6

	const (
		_   = 1 << (10 * iota) // x << n は 2のn乗による乗算
		KiB                    // 1024
		MiB                    // 1048576
		GiB
		TiB
		PiB
		EiB
		ZiB
		YiB
	)
	fmt.Println(KiB, MiB) // 1024 1048576

	// 型付けなし定数(プログラミング言語Go P87)
	// Goの定数は名前付き基本型を含めて基本データ型のどの方でも持てるが、多くの定数は特定の型に結びついているのではない。
	// コンパイラは型に結びついていない定数を基本型の値よりもはるかに高い数値制度で表現する。
	// 型に結びついていない定数は6種類。型付けなしブーリアン、型付けなし整数、型付けなしルーン、型付けなし浮動小数点数、型付けなし複素数、型付けなし文字列
	// 型への結びつけを遅延させることで、後まで高い精度を維持でき、またより多くの式で変換を必要とせずに使える。

	// YiBとZiBの値は大きすぎてどのような整数変数へも保存できないが、次のような式で使える正当な定数
	//yib := 1208925819614629174706176	// overflows int
	fmt.Println(YiB / ZiB)	// 1024

	// 変換を必要とせずに使える
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)
}
