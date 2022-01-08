package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

//ダイジェストを参照するためのプリント関数
func PrintMap(m []map[int]int) {
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, c := nextInt(), nextInt()
	d := make([][]int, c)
	for i := 0; i < c; i++ {
		d[i] = nextIntSlice(c)
	}
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = nextIntSlice(n)
		for j := 0; j < n; j++ {
			//予めマス目の色を0-indexedにしておく
			f[i][j]--
		}
	}
	m := make([]map[int]int, 3)
	for i := 0; i < 3; i++ {
		m[i] = make(map[int]int)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[(i+j)%3][f[i][j]]++
		}
	}
	//PrintMap(m)
	ans := 1 << 60
	for i := 0; i < c; i++ {
		for j := 0; j < c; j++ {
			for k := 0; k < c; k++ {
				if i == j || j == k || i == k {
					continue
				}
				s := 0
				// (i+j)%3==0のマスをiに塗り替える
				for key, v := range m[0] {
					if key == i {
						continue
					}
					s += d[key][i] * v
				}
				for key, v := range m[1] {
					if key == j {
						continue
					}
					s += d[key][j] * v
				}
				for key, v := range m[2] {
					if key == k {
						continue
					}
					s += d[key][k] * v
				}
				ans = Min(ans, s)
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

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
