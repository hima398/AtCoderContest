package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n int, x, y []int) float64 {
	ans := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i == j || j == k || k == i {
					continue
				}
				ax, ay := float64(x[i]-x[j]), float64(y[i]-y[j])
				bx, by := float64(x[k]-x[j]), float64(y[k]-y[j])
				cos := ax*bx + ay*by
				cos /= math.Sqrt(ax*ax+ay*ay) * math.Sqrt(bx*bx+by*by)
				rad := math.Acos(cos)
				ans = math.Max(ans, rad)
			}
		}
	}
	ans = ans * 180.0 / math.Pi
	return ans
}

func Solve(n int, x, y []int) (ans float64) {

	for i := 0; i < n; i++ {
		xd, yd := make([]float64, n), make([]float64, n)
		var rs []float64
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			xd[j], yd[j] = float64(x[j]-x[i]), float64(y[j]-y[i])
			rad := math.Atan2(yd[j], xd[j])
			rs = append(rs, rad, rad+2*math.Pi)
		}
		sort.Float64s(rs)

		//fmt.Println(rs)

		for j := 0; j < len(rs); j++ {
			target := rs[j] + math.Pi

			//fmt.Println("target = ", target)
			idx := sort.Search(len(rs), func(k int) bool {
				return target < rs[k]
			})
			//fmt.Println("idx = ", idx)

			rad := rs[idx-1] - rs[j]
			ans = math.Max(ans, rad)
		}
	}
	//rad -> digに変換
	ans = ans * 180.0 / math.Pi
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextInt(), nextInt()
	}

	//fmt.Println(SolveHonestly(n, x, y))
	fmt.Println(Solve(n, x, y))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
