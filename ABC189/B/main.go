package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintTestCase(n, x int, v, p []int) {
	fmt.Println(n, x)
	for i := 0; i < n; i++ {
		fmt.Println(v[i], p[i])
	}
	fmt.Println("")
}

func TestFloatingPoint() {
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < 1000000; j++ {
		n := rand.Intn(1000)
		x := rand.Intn(1000000)
		//fmt.Println(n, x)
		v := make([]int, n)
		p := make([]int, n)
		for i := 0; i < n; i++ {
			v[i] = rand.Intn(999) + 1
			p[i] = rand.Intn(100)
		}
		// WrongAnswer
		wa := 0.0
		wans := 0
		for i := 0; i < n; i++ {
			v, p := float64(v[i]), float64(p[i])
			wa += v * p / 100
			//fmt.Println(a, x)
			if wa > float64(x) {
				wans = i + 1
				break
			}
		}
		// Solve
		a := 0
		x *= 100
		ans := 0
		for i := 0; i < n; i++ {
			a += v[i] * p[i]
			if a > x {
				ans = i + 1
				break
			}
		}
		if wans != ans {
			PrintTestCase(n, x, v, p)
		}
	}
}

func WrongAnswer() {
	n, x := nextInt(), float64(nextInt())
	a := 0.0
	for i := 0; i < n; i++ {
		v, p := float64(nextInt()), float64(nextInt())
		a += v * p / 100
		//fmt.Println(a, x)
		if a > x {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}

func Solve() {
	n, x := nextInt(), nextInt()
	a := 0
	x *= 100
	for i := 0; i < n; i++ {
		v, p := nextInt(), nextInt()
		a += v * p
		//fmt.Println(a, x)
		if a > x {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	//TestFloatingPoint()
	WrongAnswer()
	//Solve()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
