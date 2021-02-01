package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func ModPow(x, y int) int {
	if y == 1 {
		return x % Mod
	}
	return (x * ModPow(x, y-1)) % Mod
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ans := ModPow(10, n) + ModPow(8, n) - (2 * ModPow(9, n) % Mod)
	// 2 * ModPow(9, n)の方が大きい場合があり
	// ansが負の数になる可能性がある。そのため正の値にする
	fmt.Println((ans + Mod) % Mod)
}
