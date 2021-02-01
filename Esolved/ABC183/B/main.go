package main

import "fmt"

func main() {
	var sx, sy, gx, gy float64
	fmt.Scan(&sx, &sy, &gx, &gy)
	x := (sy*gx + sx*gy) / (gy + sy)
	fmt.Println(x)
}
