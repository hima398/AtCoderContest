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
	var x []string
	// 番兵
	x = append(x, ".........")
	for i := 0; i < n; i++ {
		s := nextString()
		x = append(x, s)
	}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 0; j < 9; j++ {
			if x[i][j] == 'x' {
				ans++
			} else if x[i][j] == 'o' {
				if x[i-1][j] != 'o' {
					ans++
				}
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}
