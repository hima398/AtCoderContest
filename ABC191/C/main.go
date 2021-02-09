package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var di = []int{0, 0, 1, 1}
var dj = []int{0, 1, 0, 1}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	h, w := nextInt(), nextInt()
	g := make([]string, h)
	for i := 0; i < h; i++ {
		g[i] = nextString()
	}
	ans := 0
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			black := 0
			for k := 0; k < 4; k++ {
				if g[i+di[k]][j+dj[k]] == '#' {
					black++
				}
			}
			if black == 1 || black == 3 {
				ans++
			}
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
