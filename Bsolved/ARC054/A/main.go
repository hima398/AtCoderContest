package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	l, x, y, s, d := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	// 時計回りで進む距離d1、かかる時間t1
	d1 := d - s
	if d1 < 0 {
		d1 += l
	}
	t1 := float64(d1) / float64(x+y)

	// 反時計回りで進む距離d2、かかる時間t2
	d2 := s - d
	if d2 < 0 {
		d2 = l + d2
	}
	// y-x = 0があり得るのでテストケース確認
	t2 := float64(d2) / float64(y-x)
	//fmt.Println(d1, t1, d2, t2)

	var ans float64
	if t2 > 0 {
		ans = math.Min(t1, t2)
	} else {
		// x > yの時反時計回りに進めないので、時計回りに決まる
		ans = t1
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
