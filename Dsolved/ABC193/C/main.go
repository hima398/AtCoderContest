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
	m := make(map[int]bool)
	for a := 2; a*a <= n; a++ {
		aa := a
		for {
			// over flow?
			if aa*a/a != aa {
				break
			}
			aa *= a
			if aa > n {
				break
			}
			m[aa] = true
		}
	}
	//fmt.Println(m)
	fmt.Println(n - len(m))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
