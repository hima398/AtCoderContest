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

	const p = int(1e9) + 7
	h, w, k := nextInt(), nextInt(), nextInt()
	dp := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		dp[i] = make([]int, w)
	}

	dp[0][0] = 1
	mask := (1 << (w - 1)) - 1
	//bitsのビットが0の部分を横線無し、1の部分を横線ありとして
	//あみだくじとして成立しているかをチェックする
	IsAmidakuji := func(bits int) bool {
		for i := 0; i < w-2; i++ {
			l := (bits>>i)&1 == 1
			r := (bits>>(i+1))&1 == 1
			// 左右に横線があることはあみだくじとして成立しないのでfalse
			if l && r {
				return false
			}
		}
		return true
	}
	for i := 1; i <= h; i++ {
		for j := 0; j < w; j++ {
			for bits := 0; bits <= mask; bits++ {
				if IsAmidakuji(bits) {
					if (bits>>j)&1 == 1 {
						dp[i][j+1] += dp[i-1][j]
						dp[i][j+1] %= p
					} else if j > 0 && (bits>>(j-1))&1 == 1 {
						dp[i][j-1] += dp[i-1][j]
						dp[i][j-1] %= p
					} else {
						dp[i][j] += dp[i-1][j]
						dp[i][j] %= p
					}
				}
			}
		}
		//高さiについて、1以外からスタートしてk以外に辿り着く数
	}
	ans := dp[h][k-1]
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
