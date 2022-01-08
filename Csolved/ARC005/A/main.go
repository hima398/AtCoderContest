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
	var ans int
	for i := 0; i < n; i++ {
		w := nextString()
		if i == n-1 {
			//.を取り除く
			w = w[:len(w)-1]
		}
		if w == "TAKAHASHIKUN" || w == "Takahashikun" || w == "takahashikun" {
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}
