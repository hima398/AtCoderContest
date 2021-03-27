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

	n := nextInt()
	t, a := nextInt(), nextInt()
	//h := make([]int, n)
	diff := 1 << 60
	ans := -1
	for i := 0; i < n; i++ {
		h := nextInt()
		d := Abs(1000*a - 1000*t + h*6)
		if d < diff {
			diff = d
			ans = i + 1
		}
	}
	fmt.Println(ans)
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
