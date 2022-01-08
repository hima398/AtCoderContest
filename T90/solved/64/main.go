package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}
	b := make([]int, n)
	s := 0
	for i := 1; i < n; i++ {
		b[i] = a[i+1] - a[i]
		s += Abs(b[i])
	}
	b = append(b, 0)
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(s)
	ans := s
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		l, r, v := nextInt(), nextInt(), nextInt()
		before := Abs(b[l-1]) + Abs(b[r])
		if l > 1 {
			b[l-1] += v
		}
		if r < n {
			b[r] -= v
		}
		after := Abs(b[l-1]) + Abs(b[r])
		ans += (after - before)
		fmt.Fprintln(out, ans)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
