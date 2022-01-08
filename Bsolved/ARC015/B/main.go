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

	ans := make([]int, 6)
	n := nextInt()
	Eval := func(lmt, smt float64) {
		if lmt >= 35.0 {
			// 猛暑日
			ans[0]++
		}
		if lmt >= 30.0 && lmt < 35.0 {
			// 真夏日
			ans[1]++
		}
		if lmt >= 25.0 && lmt < 30.0 {
			// 夏日
			ans[2]++
		}
		if smt >= 25.0 {
			// 熱帯夜
			ans[3]++
		}
		if smt < 0.0 && lmt >= 0.0 {
			ans[4]++
		}
		if lmt < 0.0 {
			ans[5]++
		}
	}

	for i := 0; i < n; i++ {
		t1, t2 := nextFloat64(), nextFloat64()
		Eval(t1, t2)
	}

	fmt.Printf("%d", ans[0])
	for i := 1; i < 6; i++ {
		fmt.Printf(" %d", ans[i])
	}
	fmt.Println("")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}
