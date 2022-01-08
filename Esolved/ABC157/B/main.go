package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const m = 3

var a [m][m]int
var bingo [m][m]bool

func IsBingo() bool {
	var ok bool

	for i := 0; i < m; i++ {
		ok = bingo[i][0] && bingo[i][1] && bingo[i][2]
		if ok {
			return ok
		}
	}
	for j := 0; j < m; j++ {
		ok = bingo[0][j] && bingo[1][j] && bingo[2][j]
		if ok {
			return ok
		}
	}
	ok = bingo[0][0] && bingo[1][1] && bingo[2][2]
	if ok {
		return ok
	}
	return bingo[0][2] && bingo[1][1] && bingo[2][0]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			a[i][j] = nextInt()
		}
	}
	n := nextInt()
	for k := 0; k < n; k++ {
		b := nextInt()
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				if a[i][j] == b {
					bingo[i][j] = true
					if IsBingo() {
						fmt.Println("Yes")
						return
					}
				}
			}
		}
	}
	fmt.Println("No")
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

func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
