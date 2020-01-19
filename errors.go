package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.Newは基本的なerror型の値を作る関数である。引数はエラーメッセージを表す。
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// Errorメソッドを実装すれば、自分で作った型をerrorとして扱える(errorインターフェースを満たすため)
// エラーが引数エラーであることを明示する型を作る
type argError struct {
	arg int
	prob string
}
func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{2, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		}else{
			fmt.Println("f1 work", r)
		}
	}
	for _, i2 := range []int{2, 42} {
		if r2, e2 := f2(i2); e2 != nil {
			fmt.Println("f2 failed", e2)
		}else{
			fmt.Println("f2 work", r2)
		}
	}

	// 自作したエラーのデータをプログラム中で使う時には、型アサーションを使って自作したエラー型のインスタンスを作る必要がある
	// f2の2番目の返り値はerror型であり、argErrorのフィールドにアクセスできないため
	// interface型で引数として受けたり、返り値として渡されたりした値は、元の型の情報が欠落するため元の型の値を操作できないので、
	// 型アサーションを使って取得する
	_, e3 := f2(42)
	// fmt.Println(e3.arg)	// e3.arg undefined (type error has no field or method arg)
	//fmt.Println(e3.prob)  // e3.prob undefined (type error has no field or method prob)

	// 型アサーションは <変数>.(<型>) の構文となる
	// 返り値の1番目は型アサーション成功時の実際の値、2番目は成功の有無(true/false)
	if ae, ok := e3.(argError); ok {
		fmt.Println("error is argError")
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}