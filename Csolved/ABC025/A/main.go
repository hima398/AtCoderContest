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

	s := nextString()
	n := nextInt()
	ii := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			ss := string(s[i]) + string(s[j])
			if ii == n-1 {
				fmt.Println(ss)
				return
			}
			ii++
		}
	}
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
