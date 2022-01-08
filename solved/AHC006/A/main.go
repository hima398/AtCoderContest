package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const INF = 1 << 62
const N = 1000

type Pos struct {
	x, y int
}

func ComputeDist(x1, y1, x2, y2 int) int {
	return Abs(x2-x1) + Abs(y2-y1)
}

func FirstSolve(a, b, c, d []int) (m int, r []int, n int, p []Pos) {
	m = 50
	n = 102
	p = append(p, Pos{400, 400})

	type Dist struct {
		i, d int
	}
	var ds []Dist
	for i := 0; i < N; i++ {
		ds = append(ds, Dist{i, ComputeDist(400, 400, a[i], b[i]) + ComputeDist(400, 400, c[i], d[i])})
	}
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].d < ds[j].d
	})
	for i := 0; i < m; i++ {
		r = append(r, ds[i].i+1)
	}
	for i := 0; i < m; i++ {
		p = append(p, Pos{a[ds[i].i], b[ds[i].i]})
		p = append(p, Pos{c[ds[i].i], d[ds[i].i]})
	}
	p = append(p, Pos{400, 400})
	return m, r, n, p
}

func SolveHonestly(a, b, c, d []int) (m int, r []int, n int, p []Pos) {
	m = 50
	n = 102
	p = append(p, Pos{400, 400})
	for i := 0; i < m; i++ {
		r = append(r, i+1)
		p = append(p, Pos{a[i], b[i]})
		p = append(p, Pos{c[i], d[i]})
	}
	p = append(p, Pos{400, 400})
	return m, r, n, p
}

type Dist struct {
	idx, dist int
}

func PickUpByDistFromCenter(a, b, c, d []int) []Dist {
	// なるべく中心に近いところ50件をピックアップ
	var ds []Dist
	for i := 0; i < N; i++ {
		//ds = append(ds, Dist{i, ComputeDist(400, 400, a[i], b[i]) + ComputeDist(400, 400, c[i], d[i])})
		ds = append(ds, Dist{i, Max(ComputeDist(400, 400, a[i], b[i]), ComputeDist(400, 400, c[i], d[i]))})
	}
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].dist < ds[j].dist
	})
	return ds
}

func PickUpByDistBetweenStartEnd(a, b, c, d []int) []Dist {
	var ds []Dist
	for i := 0; i < N; i++ {
		ds = append(ds, Dist{i, ComputeDist(a[i], b[i], c[i], d[i])})
	}
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].dist < ds[j].dist
	})
	return ds

}

func SecondSolve(a, b, c, d []int) (m int, r []int, n int, p []Pos) {
	m = 50
	n = 102

	ds := PickUpByDistFromCenter(a, b, c, d)
	//ds := PickUpByDistBetweenStartEnd(a, b, c, d)
	// 注文を担当する部分集合を決める
	for i := 0; i < m; i++ {
		r = append(r, ds[i].idx+1)
	}

	// ルートの検索
	type Point struct {
		i, t, d, x, y int
	}
	var ps []Point
	// 中心からの距離がトータルで近い順に候補を立てる
	for i := 0; i < m; i++ {
		idx := ds[i].idx
		ps = append(ps, Point{idx, 0, 0, a[idx], b[idx]})
		ps = append(ps, Point{idx, 1, 0, c[idx], d[idx]})
	}
	var q []Point
	q = append(q, Point{-1, -1, 0, 400, 400})
	pickUp := make([]bool, N)
	for len(q) > 0 {
		point := q[0]
		q = q[1:]
		p = append(p, Pos{point.x, point.y})
		for i := range ps {
			ps[i].d = ComputeDist(point.x, point.y, ps[i].x, ps[i].y)
		}
		sort.Slice(ps, func(i, j int) bool {
			return ps[i].d < ps[j].d
		})
		for len(ps) > 0 {
			next := ps[0]
			ps = ps[1:]
			if next.t == 0 {
				pickUp[next.i] = true
				q = append(q, next)
				break
			} else if pickUp[next.i] {
				q = append(q, next)
				break
			} else {
				// まだ注文をピックアップしていない行き先なので後回しにする
				ps = append(ps, next)
			}
		}
	}

	p = append(p, Pos{400, 400})
	return m, r, n, p
}

func ThirdSolve(a, b, c, d []int) (m int, r []int, n int, p []Pos, dist int) {
	m = 50
	n = 102

	// なるべく中心に近いところ50件をピックアップ
	ds := PickUpByDistFromCenter(a, b, c, d)

	// 注文を担当する部分集合を決める
	for i := 0; i < m; i++ {
		r = append(r, ds[i].idx+1)
	}

	// ルートの検索
	type Point struct {
		i, t, d, x, y int
	}
	var starts []Point
	var ps []Point
	// 中心からの距離がトータルで近い順に候補を立てる
	for i := 0; i < m; i++ {
		idx := ds[i].idx
		starts = append(starts, Point{idx, 0, 0, a[idx], b[idx]})
		ps = append(ps, Point{idx, 0, 0, a[idx], b[idx]})
		ps = append(ps, Point{idx, 1, 0, c[idx], d[idx]})
	}

	minD := 1 << 62
	for i := 0; i < len(starts); i++ {
		var sd int
		var rr []int
		var pp []Pos
		var q []Point
		psd := make([]Point, len(ps))
		copy(psd, ps)

		q = append(q, Point{-1, -1, ComputeDist(400, 400, starts[i].x, starts[i].y), starts[i].x, starts[i].y})
		pickUp := make([]bool, N)
		pickUp[starts[i].i] = true
		for len(q) > 0 && len(rr) < m {
			point := q[0]
			q = q[1:]
			sd += point.d
			pp = append(pp, Pos{point.x, point.y})
			for j := range psd {
				psd[j].d = ComputeDist(point.x, point.y, psd[j].x, psd[j].y)
			}
			sort.Slice(psd, func(ii, jj int) bool {
				return psd[ii].d < psd[jj].d
			})
			for len(psd) > 0 {
				next := psd[0]
				psd = psd[1:]
				if next.t == 0 {
					// スタート地点と同じものは無視する
					if next.x == starts[i].x && next.y == starts[i].y {
						continue
					}
					pickUp[next.i] = true
					q = append(q, next)
					break
				} else if pickUp[next.i] {
					q = append(q, next)
					break
				} else {
					// まだ注文をピックアップしていない行き先なので後回しにする
					psd = append(psd, next)
				}
			}
		}
		sd += ComputeDist(pp[len(pp)-1].x, pp[len(pp)-1].y, 400, 400)
		//fmt.Println(i, minD, sd)
		if minD > sd {
			p = make([]Pos, len(pp))
			copy(p, pp)
			minD = sd
		}
	}
	p = append([]Pos{Pos{400, 400}}, p...)
	p = append(p, Pos{400, 400})
	return m, r, n, p, minD
}

func PickUpByArea(a, b, c, d []int) []Dist {
	area := make([][]Dist, 4)
	var other []Dist
	for i := 0; i < N; i++ {
		if a[i] <= 400 && b[i] <= 400 && c[i] <= 400 && d[i] <= 400 {
			area[0] = append(area[0], Dist{i, ComputeDist(a[i], b[i], 400, 400) + ComputeDist(c[i], d[i], 400, 400)})
			//area[0] = append(area[0], Dist{i, Max(ComputeDist(a[i], b[i], 400, 400), ComputeDist(c[i], d[i], 400, 400))})
		} else if a[i] > 400 && b[i] <= 400 && c[i] > 400 && d[i] <= 400 {
			area[1] = append(area[1], Dist{i, ComputeDist(a[i], b[i], 400, 400) + ComputeDist(c[i], d[i], 400, 400)})
			//area[1] = append(area[1], Dist{i, Max(ComputeDist(a[i], b[i], 400, 400), ComputeDist(c[i], d[i], 400, 400))})
		} else if a[i] <= 400 && b[i] > 400 && c[i] <= 400 && d[i] > 400 {
			area[2] = append(area[2], Dist{i, ComputeDist(a[i], b[i], 400, 400) + ComputeDist(c[i], d[i], 400, 400)})
			//area[2] = append(area[2], Dist{i, Max(ComputeDist(a[i], b[i], 400, 400), ComputeDist(c[i], d[i], 400, 400))})
		} else if a[i] > 400 && b[i] > 400 && c[i] > 400 && d[i] > 400 {
			area[3] = append(area[3], Dist{i, ComputeDist(a[i], b[i], 400, 400) + ComputeDist(c[i], d[i], 400, 400)})
			//area[3] = append(area[3], Dist{i, Max(ComputeDist(a[i], b[i], 400, 400), ComputeDist(c[i], d[i], 400, 400))})
		} else {
			other = append(other, Dist{i, ComputeDist(a[i], b[i], 400, 400) + ComputeDist(c[i], d[i], 400, 400)})
		}
	}
	sort.Slice(area, func(i, j int) bool {
		return len(area[i]) > len(area[j])
	})
	var ret []Dist
	for i := 0; i < 4; i++ {
		sort.Slice(area[i], func(ii, jj int) bool {
			return area[i][ii].dist < area[i][jj].dist
		})
		for j := 0; j < len(area[i]); j++ {
			ret = append(ret, area[i][j])
		}
	}

	ret = append(ret, other...)
	return ret
}

func Solve4(a, b, c, d []int) (m int, r []int, n int, p []Pos, dist int) {
	m = 50
	n = 102

	// なるべく中心に近いところ50件をピックアップ
	ds := PickUpByArea(a, b, c, d)

	// 注文を担当する部分集合を決める
	for i := 0; i < m; i++ {
		r = append(r, ds[i].idx+1)
	}

	// ルートの検索
	type Point struct {
		i, t, d, x, y int
	}
	var starts []Point
	var ps []Point
	// 中心からの距離がトータルで近い順に候補を立てる
	for i := 0; i < m; i++ {
		idx := ds[i].idx
		starts = append(starts, Point{idx, 0, 0, a[idx], b[idx]})
		ps = append(ps, Point{idx, 0, 0, a[idx], b[idx]})
		ps = append(ps, Point{idx, 1, 0, c[idx], d[idx]})
	}

	minD := 1 << 62
	for i := 0; i < len(starts); i++ {
		var sd int
		var rr []int
		var pp []Pos
		var q []Point
		psd := make([]Point, len(ps))
		copy(psd, ps)

		q = append(q, Point{-1, -1, ComputeDist(400, 400, starts[i].x, starts[i].y), starts[i].x, starts[i].y})
		pickUp := make([]bool, N)
		pickUp[starts[i].i] = true
		for len(q) > 0 && len(rr) < m {
			point := q[0]
			q = q[1:]
			sd += point.d
			pp = append(pp, Pos{point.x, point.y})
			for j := range psd {
				psd[j].d = ComputeDist(point.x, point.y, psd[j].x, psd[j].y)
			}
			sort.Slice(psd, func(ii, jj int) bool {
				return psd[ii].d < psd[jj].d
			})
			for len(psd) > 0 {
				next := psd[0]
				psd = psd[1:]
				if next.t == 0 {
					// スタート地点と同じものは無視する
					if next.x == starts[i].x && next.y == starts[i].y {
						continue
					}
					pickUp[next.i] = true
					q = append(q, next)
					break
				} else if pickUp[next.i] {
					q = append(q, next)
					break
				} else {
					// まだ注文をピックアップしていない行き先なので後回しにする
					psd = append(psd, next)
				}
			}
		}
		sd += ComputeDist(pp[len(pp)-1].x, pp[len(pp)-1].y, 400, 400)
		//fmt.Println(i, minD, sd)
		if minD > sd {
			p = make([]Pos, len(pp))
			copy(p, pp)
			minD = sd
		}
	}
	p = append([]Pos{Pos{400, 400}}, p...)
	p = append(p, Pos{400, 400})
	return m, r, n, p, minD
}

func PickUpByArea2(a, b, c, d []int) [][]Dist {
	area := make([][]Dist, 5)
	for i := 0; i < N; i++ {
		if a[i] <= 400 && b[i] <= 400 && c[i] <= 400 && d[i] <= 400 {
			//area[0] = append(area[0], Dist{i, ComputeDist(a[i], b[i], 400, 400) + ComputeDist(c[i], d[i], 400, 400)})
			area[0] = append(area[0], Dist{i, Max(ComputeDist(a[i], b[i], 400, 400), ComputeDist(c[i], d[i], 400, 400))})
		} else if a[i] > 400 && b[i] <= 400 && c[i] > 400 && d[i] <= 400 {
			//area[1] = append(area[1], Dist{i, ComputeDist(a[i], b[i], 400, 400) + ComputeDist(c[i], d[i], 400, 400)})
			area[1] = append(area[1], Dist{i, Max(ComputeDist(a[i], b[i], 400, 400), ComputeDist(c[i], d[i], 400, 400))})
		} else if a[i] <= 400 && b[i] > 400 && c[i] <= 400 && d[i] > 400 {
			//area[2] = append(area[2], Dist{i, ComputeDist(a[i], b[i], 400, 400) + ComputeDist(c[i], d[i], 400, 400)})
			area[2] = append(area[2], Dist{i, Max(ComputeDist(a[i], b[i], 400, 400), ComputeDist(c[i], d[i], 400, 400))})
		} else if a[i] > 400 && b[i] > 400 && c[i] > 400 && d[i] > 400 {
			//area[3] = append(area[3], Dist{i, ComputeDist(a[i], b[i], 400, 400) + ComputeDist(c[i], d[i], 400, 400)})
			area[3] = append(area[3], Dist{i, Max(ComputeDist(a[i], b[i], 400, 400), ComputeDist(c[i], d[i], 400, 400))})
		}
		// idx=4は全部入れておく
		area[4] = append(area[4], Dist{i, ComputeDist(a[i], b[i], 400, 400) + ComputeDist(c[i], d[i], 400, 400)})
	}
	//sort.Slice(area, func(i, j int) bool {
	//	return len(area[i]) > len(area[j])
	//})
	for i := 0; i < 5; i++ {
		sort.Slice(area[i], func(ii, jj int) bool {
			return area[i][ii].dist < area[i][jj].dist
		})
	}
	return area
}

func Solve5(a, b, c, d []int) (m int, r []int, n int, p []Pos, dist int) {
	m = 50
	n = 102

	// なるべく中心に近いところ50件をピックアップ
	ds := PickUpByArea2(a, b, c, d)

	minD := 1 << 62

	for k := 0; k < 5; k++ {
		if len(ds[k]) < 50 {
			continue
		}
		// ds[k]は50以上
		var rr []int
		// 注文を担当する部分集合を決める
		for i := 0; i < m; i++ {
			rr = append(rr, ds[k][i].idx+1)
		}

		// ルートの検索
		type Point struct {
			i, t, d, x, y int
		}
		var ps []Point
		// 中心からの距離がトータルで近い順に候補を立てる
		for i := 0; i < m; i++ {
			idx := ds[k][i].idx
			ps = append(ps, Point{idx, 0, 0, a[idx], b[idx]})
			ps = append(ps, Point{idx, 1, 0, c[idx], d[idx]})
		}

		var sd int
		var pp []Pos
		var q []Point
		psd := make([]Point, len(ps))
		copy(psd, ps)

		//q = append(q, Point{-1, -1, ComputeDist(400, 400, starts[i].x, starts[i].y), starts[i].x, starts[i].y})
		q = append(q, Point{-1, -1, 0, 400, 400})

		pickUp := make([]bool, N)
		//pickUp[starts[i].i] = true
		for len(q) > 0 {
			point := q[0]
			q = q[1:]
			sd += point.d
			pp = append(pp, Pos{point.x, point.y})
			for j := range psd {
				psd[j].d = ComputeDist(point.x, point.y, psd[j].x, psd[j].y)
			}
			sort.Slice(psd, func(ii, jj int) bool {
				return psd[ii].d < psd[jj].d
			})
			for len(psd) > 0 {
				next := psd[0]
				psd = psd[1:]
				if next.t == 0 {

					pickUp[next.i] = true
					q = append(q, next)
					break
				} else if pickUp[next.i] {
					q = append(q, next)
					break
				} else {
					// まだ注文をピックアップしていない行き先なので後回しにする
					psd = append(psd, next)
				}
			}
		}
		sd += ComputeDist(pp[len(pp)-1].x, pp[len(pp)-1].y, 400, 400)
		//fmt.Println(i, minD, sd)
		if minD > sd {
			r = make([]int, m)
			copy(r, rr)
			p = make([]Pos, len(pp))
			copy(p, pp)
			minD = sd
		}

	}
	p = append(p, Pos{400, 400})
	return m, r, n, p, minD
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	//入力
	a, b, c, d := make([]int, N), make([]int, N), make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i], c[i], d[i] = nextInt(), nextInt(), nextInt(), nextInt()
	}

	//計算
	//m, r, n, p := FirstSolve(a, b, c, d)
	//m, r, n, p := SecondSolve(a, b, c, d)
	m, r, n, p, dist := ThirdSolve(a, b, c, d)
	m4, r4, n4, p4, d4 := Solve4(a, b, c, d)
	m5, r5, n5, p5, d5 := Solve5(a, b, c, d)
	//var m, n int
	//var r []int
	//var p []Pos
	if dist > d4 {
		m, r, n, p, dist = m4, r4, n4, p4, d4
	}
	if dist > d5 {
		m, r, n, p = m5, r5, n5, p5
	}

	//出力
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintf(out, "%d", m)
	for i := range r {
		fmt.Fprintf(out, " %d", r[i])
	}
	fmt.Fprintln(out)
	fmt.Fprintf(out, "%d", n)
	for i := range p {
		fmt.Fprintf(out, " %d %d", p[i].x, p[i].y)
	}
	//fmt.Println(len(r), len(p))
	fmt.Fprintln(out)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	/*
		if x < y {
			x, y = y, x
		}
	*/
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	// x*yのオーバーフロー対策のため先にGcdで割る
	// Gcd(x, y)はxの約数のため割り切れる
	ret := x / Gcd(x, y)
	ret *= y
	return ret
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}

func Permutation(N, K int) int {
	v := 1
	if 0 < K && K <= N {
		for i := 0; i < K; i++ {
			v *= (N - i)
		}
	} else if K > N {
		v = 0
	}
	return v
}

func Factional(N int) int {
	return Permutation(N, N-1)
}

func Combination(N, K int) int {
	if K == 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}

type Comb struct {
	n, p int
	fac  []int // Factional(i) mod p
	finv []int // 1/Factional(i) mod p
	inv  []int // 1/i mod p
}

func NewCombination(n, p int) *Comb {
	c := new(Comb)
	c.n = n
	c.p = p
	c.fac = make([]int, n+1)
	c.finv = make([]int, n+1)
	c.inv = make([]int, n+1)

	c.fac[0] = 1
	c.fac[1] = 1
	c.finv[0] = 1
	c.finv[1] = 1
	c.inv[1] = 1
	for i := 2; i <= n; i++ {
		c.fac[i] = c.fac[i-1] * i % p
		c.inv[i] = p - c.inv[p%i]*(p/i)%p
		c.finv[i] = c.finv[i-1] * c.inv[i] % p
	}
	return c
}

func (c *Comb) Factional(x int) int {
	return c.fac[x]
}

func (c *Comb) Combination(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	ret := c.fac[n] * c.finv[k]
	ret %= c.p
	ret *= c.finv[n-k]
	ret %= c.p
	return ret
}

//重複組み合わせ H
func (c *Comb) DuplicateCombination(n, k int) int {
	return c.Combination(n+k-1, k)
}
func (c *Comb) Inv(x int) int {
	return c.inv[x]
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func DivideSlice(A []int, K int) ([]int, []int, error) {

	if len(A) < K {
		return nil, nil, errors.New("")
	}
	return A[:K+1], A[K:], nil
}

type IntQueue struct {
	q []int
}

func NewIntQueue() *IntQueue {

	return new(IntQueue)
}
func (this *IntQueue) Push(v int) {
	this.q = append(this.q, v)
}

func (this *IntQueue) Pop() (int, error) {
	if this.Size() == 0 {
		return -1, errors.New("")
	}
	ret := this.q[0]
	this.q = this.q[1:]
	return ret, nil
}

func (this *IntQueue) Size() int {
	return len(this.q)
}

func (this *IntQueue) PrintQueue() {
	fmt.Println(this.q)
}
