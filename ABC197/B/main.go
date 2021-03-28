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

	h, w, x, y := nextInt(), nextInt(), nextInt(), nextInt()
	x--
	y--
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	ans := 0
	for i := x; i >= 0; i-- {
		if s[i][y] == '.' {
			//fmt.Println(i, y)
			ans++
		} else {
			break
		}
	}
	for i := x; i < h; i++ {
		if s[i][y] == '.' {
			//fmt.Println(i, y)
			ans++
		} else {
			break
		}
	}
	for j := y; j >= 0; j-- {
		if s[x][j] == '.' {
			//fmt.Println(x, j)
			ans++
		} else {
			break
		}
	}
	for j := y; j < w; j++ {
		if s[x][j] == '.' {
			//fmt.Println(x, j)
			ans++
		} else {
			break
		}
	}
	ans -= 3
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
