package main

import "fmt"

type Span struct {
	start  int
	end    int
	amount int
}

func main() {
	// 1 <= n <= 2 * 10**5
	// 1 <= w
	var n, w int
	fmt.Scan(&n, &w)
	// 0 < si < ti < 2 * 10**5
	s := make([]int, n)
	t := make([]int, n)
	// pi <= 10**9 int64?
	p := make([]int, n)
	m := make(map[int]*Span)
	for i := 0; i < n; i++ {
		if m[s[i]] == nil {
			if p[i] > w {
				fmt.Println("No")
				return
			}
			span := &Span{s[i], t[i], p[i]}
			m[span.start] = span
		} else {
			// m[si] exist
			idx := s[i]
			next := idx
			for idx < t[i] {
				m[idx].amount += p[i]
				if m[idx].amount > w {
					fmt.Println("No")
					return
				}
				next = m[idx].end + 1
				idx = m[idx].end
			}
			if next < t[i] {
				span := &Span{next, t[i], p[i]}
				m[next] = span
			}
		}
	}
	fmt.Println(m)
	fmt.Println("Yes")
}
