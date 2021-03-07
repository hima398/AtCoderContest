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
	ng := 0
	for i := 0; i < n; i++ {
		p := nextInt()
		if i+1 != p {
			ng++
		}
	}
	if ng == 0 || ng == 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
