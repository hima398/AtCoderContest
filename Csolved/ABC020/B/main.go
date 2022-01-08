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

	a, b := nextString(), nextString()
	ans, _ := strconv.Atoi(a + b)
	ans *= 2
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
