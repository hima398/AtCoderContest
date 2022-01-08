package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n int) (ans int) {
	for i := 1; i <= n; i++ {
		nn := i
		for nn > 0 {
			if nn%10 == 1 {
				ans++
			}
			nn /= 10
		}
	}
	return ans
}

func Solve(n int) (ans int) {
	for i := 1; i <= n; i *= 10 {
		ans += i * (n / i / 10)
		//fmt.Println(i, ans)
		t := n / i % 10
		if t > 1 {
			ans += i
		} else if t == 1 {
			ans += n%i + 1
		}
		//fmt.Println(i, ans)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()

	//fmt.Println(SolveHonestly(n))
	fmt.Println(Solve(n))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
