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
	//a := make([]int, n)
	m := make(map[int]bool)
	rem := 0
	for i := 0; i < n; i++ {
		ai := nextInt()
		if !m[ai] {
			m[ai] = true
		} else {
			rem++
		}
	}
	ans := len(m)
	if rem%2 == 1 {
		ans--
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
