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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextString()
	ans := 0
	for i := 0; i < len(n); i++ {
		s, _ := strconv.Atoi(string(n[i]))
		ans += s
	}
	if ans%9 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
