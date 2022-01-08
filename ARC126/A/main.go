package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve(n2, n3, n4 int) (ret int) {
	// 3本の棒を使う
	n6 := n3 / 2
	//fmt.Println("ret = ", ret)
	d := Min(n6, n4)
	ret += d
	n4 -= d
	n6 -= d

	d = Min(n6, n2/2)
	ret += d
	n2 -= 2 * d

	n8 := n4 / 2
	n4 -= 2 * n8

	d = Min(n8, n2)
	ret += d
	//fmt.Println("ret = ", ret)
	n2 -= d

	if n4 > 0 && n2 >= 3 {
		// 長さ4と長さ2 * 3をつなげる
		ret += 1
		n2 -= 3
	}

	ret += n2 / 5
	//fmt.Println("ret = ", ret)

	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	ans := make([]int, t)
	for i := 0; i < t; i++ {
		n2, n3, n4 := nextInt(), nextInt(), nextInt()
		ans[i] = FirstSolve(n2, n3, n4)
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
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
