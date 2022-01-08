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
	ans := n - 2021 // n-2021+1(2021年が6回目)
	if ans < 0 {
		switch n {
		case 2015:
			ans = 1
		case 2016:
			ans = 2
		case 2018:
			ans = 3
		case 2019:
			ans = 4
		case 2020:
			ans = 5
		default:
			ans = -1
		}
	} else {
		ans += 6
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
