package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i].X)
		fmt.Scan(&p[i].Y)
	}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			a1 := (p[j].Y - p[i].Y) / (p[j].X - p[i].X)
			for k := j + 1; k < n; k++ {
				if p[j].X == p[i].X {
					if p[k].X == p[i].X {
						fmt.Println("Yes")
						return
					}
				} else {
					a2 := (p[k].Y - p[i].Y) / (p[k].X - p[i].X)
					if a1 == a2 {
						fmt.Println("Yes")
						return
					}
				}
			}
		}
	}
	fmt.Println("No")
	return
}
