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

	r, g, b, n := nextInt(), nextInt(), nextInt(), nextInt()
	var ans int
	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			kb := n - (i*r + j*g)
			if kb >= 0 && kb%b == 0 {
				//fmt.Println(i, j, kb/b)
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
