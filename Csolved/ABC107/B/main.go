package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintGrid(s []string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	for i := 0; i < len(s); i++ {
		canCompress := true
		for j := 0; j < w; j++ {
			canCompress = canCompress && s[i][j] == '.'
		}
		if canCompress {
			s = append(s[:i], s[i+1:]...)
			//次のループでi=0から初めて欲しいのでi++されて0になるようにiを-1にする
			i = -1
		}
	}

	for j := 0; j < len(s[len(s)-1]); j++ {
		canCompress := true
		for i := 0; i < len(s); i++ {
			canCompress = canCompress && s[i][j] == '.'
		}

		if canCompress {
			for i := 0; i < len(s); i++ {
				ss := strings.Split(s[i], "")

				s[i] = strings.Join(append(ss[:j], ss[j+1:]...), "")
			}
			j = -1
		}
	}
	PrintGrid(s)
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
