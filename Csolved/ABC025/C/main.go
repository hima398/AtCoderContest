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

	var f [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			f[i][j] = -1
		}
	}
	var b [2][3]int
	var c [3][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = nextInt()
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = nextInt()
		}
	}

	var Dfs func(turn int) (int, int)
	Dfs = func(turn int) (int, int) {
		if turn == 9 {
			var nh, nk int
			for i := 0; i < 2; i++ {
				for j := 0; j < 3; j++ {
					if f[i+1][j] == f[i][j] {
						nh += b[i][j]
					} else {
						nk += b[i][j]
					}
				}
			}
			for i := 0; i < 3; i++ {
				for j := 0; j < 2; j++ {
					if f[i][j+1] == f[i][j] {
						nh += c[i][j]
					} else {
						nk += c[i][j]
					}
				}
			}
			//fmt.Println(turn, nh, nk)
			return nh, nk
		}
		nh, nk := -1, -1
		IsNaokosTurn := turn % 2
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				//o, xがまだ未記入
				if f[i][j] == -1 {
					f[i][j] = IsNaokosTurn
					s1, s2 := Dfs(turn + 1)
					if nh < 0 {
						nh, nk = s1, s2
					} else {
						//直大くん、直子さんそれぞれturn+1ターン目のスコアが大きい場合に更新
						if IsNaokosTurn == 0 {
							if nh < s1 {
								nh, nk = s1, s2
							}
						} else {
							if nk < s2 {
								nh, nk = s1, s2
							}
						}
						//それぞれturn+1ターン目のスコアが同じ場合、相手のスコアが小さい場合に更新
						if IsNaokosTurn == 0 {
							if nh == s1 && nk > s2 {
								nh, nk = s1, s2
							}
						} else {
							if nh > s1 && nk == s2 {
								nh, nk = s1, s2
							}
						}
					}
					f[i][j] = -1
				}
			}
		}
		return nh, nk
	}
	nh, nk := Dfs(0)
	fmt.Println(nh)
	fmt.Println(nk)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
