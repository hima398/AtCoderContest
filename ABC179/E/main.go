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

	n, x, m := nextInt(), nextInt(), nextInt()
	app := make([]int, m)
	var route []int
	var loop []int
	sl := 0
	for a := x; app[a] < 2; a = a * a % m {
		if app[a] == 1 {
			loop = append(loop, a)
			sl += a
		} else {
			route = append(route, a)
		}
		app[a]++
	}
	//fmt.Println(len(route))
	//fmt.Println(len(loop))
	ans := 0
	b := len(route) - len(loop)
	if b > n {
		for i := 0; i < n; i++ {
			ans += route[i]
		}
		fmt.Println(ans)
		return
	}
	for i := 0; i < b; i++ {
		ans += route[i]
	}
	n -= b
	ans += sl * (n / len(loop))
	for i := 0; i < n%len(loop); i++ {
		ans += loop[i]
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
