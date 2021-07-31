package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	const maxK = 40
	fib := []int{1, 1}
	for i := 0; i < maxK; i++ {
		n := len(fib)
		fib = append(fib, fib[n-1]+fib[n-2])
	}
	k := nextInt()

	fmt.Println(fib[k+1], fib[k])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
