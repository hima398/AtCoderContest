package main

import "fmt"

// データの入出力用テンプレート
func main() {
	var a int
	fmt.Scan(&a)

	var b, c int
	fmt.Scan(&b, &c)

	var s string
	fmt.Scan(&s)

	fmt.Printf("%d %s", a+b+c, s)
}
