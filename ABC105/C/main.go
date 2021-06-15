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
	if n == 0 {
		fmt.Println(0)
		return
	}

	var rs []int
	for n != 0 {
		r := n % 2
		if r < 0 {
			r *= -1
		}
		rs = append(rs, r)
		n = -(n - r) / 2
	}
	ans := ""
	for i := 0; i < len(rs); i++ {
		si := strconv.Itoa(rs[len(rs)-i-1])
		ans += si
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
