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

	n := nextInt()
	var a []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
			if n/i != i {
				a = append(a, n/i)
			}
		}
	}
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(a[i])
	}

}
