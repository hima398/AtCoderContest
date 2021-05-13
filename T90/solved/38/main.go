package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve(a, b int) {
	ba, bb := big.NewInt(int64(a)), big.NewInt(int64(b))
	gcd := big.NewInt(int64(Gcd(a, b)))
	max := big.NewInt(int64(1e18))
	ans := big.NewInt(1).Mul(ba, bb)
	ans = ans.Div(ans, gcd)
	if ans.Cmp(max) == 1 {
		//if ans <= 0 || ans > max || ans%a != 0 || ans%b != 0 {
		fmt.Println("Large")
		return
	}
	fmt.Println(ans)
}

func Solve(a, b int) {
	const max = int(1e18)
	// 必ずaの方がb以上になるようにする
	if a < b {
		a, b = b, a
	}
	gcd := Gcd(a, b)
	c := a / gcd
	ans := c * b
	// * b / bで元に戻るかでオーバーフローと、最大値との比較で検査する
	if ans/b != c || ans > max {
		fmt.Println("Large")
		return
	}
	fmt.Println(ans)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextInt(), nextInt()
	//FirstSolve(a, b)
	Solve(a, b)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	if x < y {
		x, y = y, x
	}
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}
