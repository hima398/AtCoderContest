package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextFloat() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

var cache [101][101][101]float64

func DP(x, y, z int) float64 {
	if cache[x][y][z] > 0 {
		return cache[x][y][z]
	}
	if x >= 100 || y >= 100 || z >= 100 {
		return 0
	}
	total := float64(x + y + z)
	ans := 0.0
	ans += (DP(x+1, y, z) + 1.0) * float64(x) / total
	ans += (DP(x, y+1, z) + 1.0) * float64(y) / total
	ans += (DP(x, y, z+1) + 1.0) * float64(z) / total
	cache[x][y][z] = ans
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, c := nextInt(), nextInt(), nextInt()

	result := DP(a, b, c)
	fmt.Printf("%v\n", result)

}
