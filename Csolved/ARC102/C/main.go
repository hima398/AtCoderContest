package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n, k int) int {
	x := n / k
	if k%2 == 1 {
		return x * x * x
	} else {
		y := (n + k/2) / k
		return x*x*x + y*y*y
	}
}

func SolveHonestly(n, k int) int {
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		r[i%k]++
	}
	ans := 0
	//min := Min(n, k)
	for a := 0; a < k; a++ {
		if a > n {
			break
		}
		b := (k - a) % k
		c := (k - a) % k
		if (b+c)%k != 0 {
			continue
		}
		ans += r[a] * r[b] * r[c]
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	//fmt.Println(SolveHonestly(n, k))
	fmt.Println(Solve(n, k))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
