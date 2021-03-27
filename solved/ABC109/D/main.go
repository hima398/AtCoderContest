package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Survey(h, w int, a [][]int) {
	sum := 0
	var odd, even int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			sum += a[i][j]
			if a[i][j]%2 == 0 {
				even++
			} else {
				odd++
			}
		}
	}
	fmt.Println(sum, odd, even)
}

type Turn struct {
	fi, fj, ti, tj int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	h, w := nextInt(), nextInt()
	a := make([][]int, h)
	//v := make([][]bool, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		//v[i] = make([]bool, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			a[i][j] = nextInt()
		}

	}
	//Survey(h, w, a)
	var ans []Turn
	isOdd := 0
	for i := 0; i < h; i++ {
		if i%2 == 0 {
			for j := 0; j < w; j++ {
				isOdd ^= a[i][j] % 2
				if isOdd == 1 {
					if j == w-1 {
						if i < h-1 {
							ans = append(ans, Turn{i + 1, j + 1, i + 2, j + 1})
						}
					} else if j < w-1 {
						ans = append(ans, Turn{i + 1, j + 1, i + 1, j + 2})
					}
				}
			}
		} else {
			for j := w - 1; j >= 0; j-- {
				isOdd ^= a[i][j] % 2
				if isOdd == 1 {
					if j == 0 {
						if i < h-1 {
							ans = append(ans, Turn{i + 1, j + 1, i + 2, j + 1})
						}
					} else if j > 0 {
						ans = append(ans, Turn{i + 1, j + 1, i + 1, j})
					}
				}
			}
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintf(out, "%d %d %d %d\n", v.fi, v.fj, v.ti, v.tj)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
