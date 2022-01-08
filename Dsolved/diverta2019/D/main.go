package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n int) (ans int) {
	for i := 1; i <= n; i++ {
		if n/i == n%i {
			ans += i
		}
	}
	return ans
}

func Solve(n int) (ans int) {
	var ds []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ds = append(ds, i)
			ds = append(ds, n/i)
		}
	}
	for _, v := range ds {
		if v > 1 && n/(v-1) == n%(v-1) {
			ans += (v - 1)
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	//fmt.Println(SolveHonestly(n))
	fmt.Println(Solve(n))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
