package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n, q int, x, y, z, w []int) int {
	const p = int(1e9) + 7

	//0-indexed
	for i := 0; i < q; i++ {
		x[i]--
		y[i]--
		z[i]--
	}
	ans := 1
	for k := 0; k < 60; k++ {
		//長さnの数列の各数値のうち、下位k bit目が条件に一致するかを全探索する
		cnt := 0
		for pat := 0; pat < 1<<n; pat++ {
			ok := true
			for i := 0; i < q; i++ {
				ok = ok && ((pat>>x[i]&1)|(pat>>y[i])&1|(pat>>z[i])&1) == w[i]>>k&1
				if !ok {
					break
				}
			}
			if ok {
				cnt++
			}
		}
		ans *= cnt
		ans %= p
	}
	return ans
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	x, y, z, w := make([]int, q), make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		x[i], y[i], z[i], w[i] = nextInt(), nextInt(), nextInt(), nextInt()
	}

	ans := Solve(n, q, x, y, z, w)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
