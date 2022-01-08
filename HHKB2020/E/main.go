package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintT(t [][][4]int, d int) {
	buf := make([][]int, len(t))
	for i := 0; i < len(t); i++ {
		buf[i] = make([]int, len(t[i]))
		for j := 0; j < len(t[i]); j++ {
			buf[i][j] = t[i][j][d]
		}
	}

	for i := 0; i < len(t); i++ {
		fmt.Println(buf[i])
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = int(1e9) + 7
	h, w := nextInt(), nextInt()
	s := make([]string, h)
	t := make([][][4]int, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
		t[i] = make([][4]int, w)
	}
	k := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '.' {
				k++
			}
		}
	}
	//fmt.Println(k)
	// 左に何マスあるか計算
	for j := 1; j < w; j++ {
		for i := 0; i < h; i++ {
			if s[i][j] == '#' || s[i][j-1] == '#' {
				continue
			}
			t[i][j][0] = t[i][j-1][0] + 1
		}
	}
	// 上
	for i := 1; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' || s[i-1][j] == '#' {
				continue
			}
			t[i][j][1] = t[i-1][j][1] + 1
		}
	}
	// 右
	for j := w - 2; j >= 0; j-- {
		for i := 0; i < h; i++ {
			if s[i][j] == '#' || s[i][j+1] == '#' {
				continue
			}
			t[i][j][2] = t[i][j+1][2] + 1
		}
	}
	// 下
	for i := h - 2; i >= 0; i-- {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' || s[i+1][j] == '#' {
				continue
			}
			t[i][j][3] = t[i+1][j][3] + 1
		}
	}
	// tの値のテスト出力
	/*
		PrintT(t, 0)
		fmt.Println("")
		PrintT(t, 1)
		fmt.Println("")
		PrintT(t, 2)
		fmt.Println("")
		PrintT(t, 3)
		fmt.Println("")
	*/
	ans := k * Pow(2, k, p)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				continue
			}
			tij := 1
			for k := 0; k < 4; k++ {
				tij += t[i][j][k]
			}
			//fmt.Println(i, j, k-tij)
			ans = ans - Pow(2, k-tij, p) + p
			/*
				if ans < 0 {
					ans += p
				}
			*/
			ans %= p
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

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}
