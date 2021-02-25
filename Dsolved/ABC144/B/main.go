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
	for i := 1; i < 10; i++ {
		if n%i == 0 {
			d := n / i
			if d < 10 {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
	return
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
