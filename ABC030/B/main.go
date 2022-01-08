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

	n, m := nextInt(), nextInt()
	if n >= 12 {
		n -= 12
	}
	he := 360.0 / 12.0
	he2 := (360.0 / 12.0) / 60.0
	me := 360.0 / 60.0

	fn := float64(n)*he + float64(m)*he2
	fm := float64(m) * me
	ans := math.Min(math.Abs(fn-fm), math.Abs(fm-fn))
	if ans > 180.0 {
		ans = 360.0 - ans
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
