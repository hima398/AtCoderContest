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

	n, m := nextInt(), nextInt()
	var ans int
	if n >= 2 && m >= 2 {
		ans = (n - 2) * (m - 2)
	} else if n == 1 && m == 1 {
		ans = 1
	} else if n == 1 {
		ans = m - 2
	} else {
		ans = n - 2
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
