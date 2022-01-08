package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		x := n / i
		ans += x
	}
	return ans
}

func SolveCommentary(n int) int {
	var ans int
	for i := 1; i <= n; {
		x := n / i
		d := n/x + 1
		ans += x * (d - i)
		i = d
	}

	return ans
}

func FirstSolve(n int) int {

	k := int(math.Floor(math.Sqrt(float64(n))))
	ans := 0
	for i := 1; i <= n/(k+1); i++ {
		ans += n / i
	}
	for i := 1; i <= k; i++ {
		ans += ((n / i) - n/(i+1)) * i
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()

	//ans = SolveHonestly(n)
	//ans := SolveCommentary(n)
	ans := FirstSolve(n)

	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
