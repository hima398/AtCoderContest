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

	x, y := nextInt(), nextInt()
	k := nextInt()

	var ans int
	// 裏を面にするだけ
	if y >= k {
		ans = x + k
	} else {
		// 面もいくつか裏にする
		ans = x - (k - y) + y
	}

	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
