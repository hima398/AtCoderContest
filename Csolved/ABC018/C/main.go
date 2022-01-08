package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintField(f [][][2]int, m int) {
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			fmt.Printf("%d ", f[i][j][m])
		}
		fmt.Println("")
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	r, c, k := nextInt(), nextInt(), nextInt()
	s := make([]string, r)
	f := make([][][2]int, r)
	for i := 0; i < r; i++ {
		f[i] = make([][2]int, c)
	}
	for i := 0; i < r; i++ {
		s[i] = nextString()
	}
	for j := 0; j < c; j++ {
		//マスi, jを含めて上にある連続するoの数
		if s[0][j] == 'o' {
			f[0][j][0] = 1
		} else {
			f[0][j][0] = 0
		}
		for i := 1; i < r; i++ {
			if s[i][j] == 'o' {
				f[i][j][0] = f[i-1][j][0] + 1
			} else {
				f[i][j][0] = 0
			}
		}
		//マスi, jを含めて下にある連続するoの数
		if s[r-1][j] == 'o' {
			f[r-1][j][1] = 1
		} else {
			f[r-1][j][1] = 0
		}
		for i := r - 2; i >= 0; i-- {
			if s[i][j] == 'o' {
				f[i][j][1] = f[i+1][j][1] + 1
			} else {
				f[i][j][1] = 0
			}
		}
	}
	//PrintField(f, 0)
	ans := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ok := true
			idx := 0
			for jj := j - k + 1; jj <= j; jj++ {
				idx++
				if jj < 0 {
					ok = false
					break
				}
				ok = ok && f[i][jj][0] >= idx && f[i][jj][1] >= idx
			}
			idx = 0
			for jj := j + k - 1; jj > j; jj-- {
				idx++
				if jj >= c {
					ok = false
					break
				}
				ok = ok && f[i][jj][0] >= idx && f[i][jj][1] >= idx
			}
			if ok {
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
