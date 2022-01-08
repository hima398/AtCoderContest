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

	h, w, k := nextInt(), nextInt(), nextInt()
	g := make([]string, h)
	for i := 0; i < h; i++ {
		g[i] = nextString()
	}
	ans := 0
	for rm := 0; rm < (1 << h); rm++ {
		for cm := 0; cm < (1 << w); cm++ {
			nb := 0
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if rm>>i&1 == 0 && cm>>j&1 == 0 && g[i][j] == '#' {
						nb++
					}
				}
			}
			if nb == k {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
