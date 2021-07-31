package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ComputeVolume(a, b, x, r, h int) float64 {
	fr, fh := float64(r), float64(h)

	ret := fr * fr * math.Pi * fh / 3.0
	if a > x {
		h1 := float64(h - (a - x))
		r1 := fr * h1 / fh
		ret = r1 * r1 * math.Pi * h1 / 3.0
	}
	if b < x+h {
		h2 := float64(h - (b - x))
		r2 := fr * h2 / fh
		ret -= r2 * r2 * math.Pi * h2 / 3.0
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	x, r, h := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], r[i], h[i] = nextInt(), nextInt(), nextInt()
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		a, b := nextInt(), nextInt()
		ans := 0.0
		for j := 0; j < n; j++ {
			if b <= x[j] || a >= x[j]+h[j] {
				continue
			}
			ans += ComputeVolume(a, b, x[j], r[j], h[j])
		}
		fmt.Fprintln(out, ans)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
