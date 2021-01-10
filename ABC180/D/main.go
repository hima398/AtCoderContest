package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	x, y, a, b := nextInt(), nextInt(), nextInt(), nextInt()

	power := x
	ans := 0
	for power < b/(a-1) && power < y/a {
		power *= a
		ans++
	}
	ans += (y - power - 1) / b
	fmt.Println(ans)
}
