package main

import (
	"fmt"
	"runtime"
)

func main() {
	slices()
}

func slices() {
	s := getSubSlice()
	fmt.Println("s", len(s), cap(s))
	printMemStat()

	all := make([][]int, 0)
	all = append(all, s)
	fmt.Println("all", len(all), cap(all))
	printMemStat()

	for i := 1; i < 10; i++ {
		s2 := getSubSlice()
		fmt.Println(i, "s2", len(s2), cap(s2))
		runtime.GC()
		printMemStat()

		all = append(all, s2)
	}

	runtime.GC()
	printMemStat()

	fmt.Println(all)
}

func getSubSlice() []int {
	s := make([]int, 1_000_000)

	return s[:1]

	//var ss = make([]int, 0, 1)
	//copy(s, ss)
	//return ss
}

func printMemStat() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc / 1024 / 1024)
}
