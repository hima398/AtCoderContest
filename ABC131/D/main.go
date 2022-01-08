package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Task struct {
	a int
	b int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var tasks []Task
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		tasks = append(tasks, Task{a, b})
	}
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].b < tasks[j].b })

	time := 0
	for _, v := range tasks {
		time += v.a
		if time > v.b {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
