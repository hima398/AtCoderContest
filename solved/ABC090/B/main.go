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

	a, b := nextInt(), nextInt()
	ans := 0
	for i := a; i <= b; i++ {
		si := strconv.Itoa(i)
		ok := true
		for j := 0; j < len(si)/2; j++ {
			ok = ok && si[j] == si[len(si)-j-1]
		}
		if ok {
			ans++
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
