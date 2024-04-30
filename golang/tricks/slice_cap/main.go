package main

import (
	"fmt"
)

type MyInt struct {
	value int
}

func main() {
	// ----------------------------------------
	s1 := make([]MyInt, 0, 2)
	s1 = append(s1, MyInt{100})
	s2 := append(s1, MyInt{200})

	ref0 := &s1[0]
	ref0.value = ref0.value + 42

	fmt.Println(s1, s2) // [{142}] [{142} {200}]

	s1 = append(s2, MyInt{111})
	ref0.value = ref0.value * 10
	fmt.Println(s1, s2) // [{142}] [{142} {200}]
	// ----------------------------------------
	i32 := make([]int32, 0)
	i32 = append(i32, []int32{1, 2, 3}...)
	fmt.Println(len(i32), cap(i32), i32) // 3 3 [1 2 3]

	i32 = append(i32, 4)

	fmt.Println(len(i32), cap(i32), i32) // 4 4 [1 2 3 4]

	i64 := make([]int64, 0)
	i64 = append(i64, []int64{1, 2, 3}...)
	fmt.Println(len(i64), cap(i64), i64) // 3 3 [1 2 3]

	i64 = append(i64, 4)

	fmt.Println(len(i64), cap(i64), i64) // 4 6 [1 2 3 4]
}
