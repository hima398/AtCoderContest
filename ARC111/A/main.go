package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveCommentary(n, m int) string {
	ans := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(n)), big.NewInt(int64(m*m)))
	ans = big.NewInt(0).Div(ans, big.NewInt(int64(m)))
	return ans.String()
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	fmt.Println(SolveCommentary(n, m))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Floor(x, y int) int {
	return x / y
}
