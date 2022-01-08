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

	m := nextInt()
	var vv int
	if m < 100 {
		vv = 0
	} else if m >= 100 && m <= 5000 {
		vv = 10 * m / 1000
	} else if m >= 6000 && m <= 30000 {
		vv = m/1000 + 50
	} else if m >= 35000 && m <= 70000 {
		vv = (m/1000-30)/5 + 80
	} else {
		vv = 89
	}
	fmt.Printf("%02d\n", vv)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
