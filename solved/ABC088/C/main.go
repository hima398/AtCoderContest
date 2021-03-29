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

	const n = 3
	var c [n][n]int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = nextInt()
		}
	}
	a := []int{0}
	b := []int{c[0][0]}
	a = append(a, c[1][0]-b[0])
	b = append(b, c[0][1]-a[0])
	a = append(a, c[2][0]-b[0])
	b = append(b, c[0][2]-a[0])
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i]+b[j] != c[i][j] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
