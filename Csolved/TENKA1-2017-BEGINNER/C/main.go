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

	ln := nextInt()
	for h := 1; h <= 3500; h++ {
		for w := 1; w <= 3500; w++ {
			d := 4*h*w - w*ln - h*ln
			if d > 0 && (h*w*ln)%d == 0 {
				n := (h * w * ln) / d
				if n > 0 {
					fmt.Println(h, n, w)
					return
				}
			}
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
