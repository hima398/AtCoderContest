package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ComputePercent(a, b int) float64 {
	fa, fb := float64(a), float64(b)
	return 100.0 * fb / (fa + fb)
}

// aグラムの水にbグラムが溶かせるか
func CanDissolved(a, b, e int) bool {
	return 100*b <= a*e
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, c, d, e, f := nextInt(), nextInt(), nextInt(), nextInt(), nextInt(), nextInt()

	//iグラムの水にjグラムの砂糖を溶かした砂糖水が作れるか
	dp := make([][]bool, 3001)
	for i := 0; i < 3001; i++ {
		dp[i] = make([]bool, 3001)
	}
	dp[0][0] = true
	for i := 0; i < 3001; i++ {
		for j := 0; j < 3001; j++ {
			if dp[i][j] {
				if i+100*a+j <= f {
					dp[i+100*a][j] = true //math.Max(dp[i+100*a][j], ComputePercent(i+100*a, j))
				}
				if i*100*b+j <= f {
					dp[i+100*b][j] = true //math.Max(dp[i+100*b][j], ComputePercent(i+100*b, j))
				}
				if i+j+c <= f && CanDissolved(i, j+c, e) {
					dp[i][j+c] = true
				}
				if i+j+d <= f && CanDissolved(i, j+d, e) {
					dp[i][j+d] = true
				}
			}
		}
	}
	//砂糖が解かせない場合でも砂糖水は最低100Aグラム以上になる
	answ, anss := 100*a, 0
	per := 0.0
	for i := 0; i < 3001; i++ {
		for j := 0; j < 3001; j++ {
			if dp[i][j] {
				nper := ComputePercent(i, j)
				if per < nper {
					answ, anss = i+j, j
					per = nper
				}
			}
		}
	}

	fmt.Println(answ, anss)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
