package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintImage(s [][]int) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	s := make([][]int, h)
	t := make([][]int, h)
	c := make([][]int, h)
	for i := 0; i < h; i++ {
		s[i] = make([]int, w)
		t[i] = make([]int, w)
		c[i] = make([]int, w)
		si := nextString()
		for j := 0; j < w; j++ {
			if si[j] == '.' {
				s[i][j] = 1
			}
		}
		copy(t[i], s[i])
	}

	di := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dj := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < 8; k++ {
				ni, nj := i+di[k], j+dj[k]
				if ni >= 0 && ni < h && nj >= 0 && nj < w {
					if s[i][j] == 1 {
						t[ni][nj] = 1
					}
				}
			}
		}
	}
	for i := 0; i < h; i++ {
		copy(c[i], t[i])
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < 8; k++ {
				ni, nj := i+di[k], j+dj[k]
				if ni >= 0 && ni < h && nj >= 0 && nj < w {
					if t[i][j] == 0 {
						c[ni][nj] = 0
					}
				}
			}
		}
	}
	IsPossible := true
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] != c[i][j] {
				IsPossible = false
				break
			}
		}
	}
	if IsPossible {
		fmt.Println("possible")
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				var c string
				if t[i][j] == 0 {
					c = "#"
				} else {
					c = "."
				}
				fmt.Printf("%s", c)
			}
			fmt.Println("")
		}
	} else {
		fmt.Println("impossible")
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
