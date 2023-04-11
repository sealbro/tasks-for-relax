package company_cases_1

import (
	"sync"
	"time"
)

func Case1() ([]int, []int, []int) {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	//println(x)
	y := append(x, 3)
	//println(x, y)
	z := append(x, 4)
	//println(x, y, z)

	return x, y, z
}

func Case2() int {
	ts := time.Now()
	_, _ = <-work(), <-work()
	return int(time.Since(ts).Seconds())
}

func work() chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		ch <- struct{}{}
	}()

	return ch
}

func Case3() []int {
	return remove([]int{2, 0, 0, 1, 0, 0, 1, 0})
}

func remove(nums []int) []int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] == 0) {
		return []int{}
	}

	j := 0

	for _, val := range nums {
		if val != 0 {
			nums[j] = val
			j++
		}
	}

	return nums[:j]
}

func Case4() int {
	size := 100
	wg := sync.WaitGroup{}
	wg.Add(size)
	mutex := makeMutex()
	counter := 0

	for i := 0; i < size; i++ {
		go func() {
			mutex.Lock()
			time.Sleep(time.Millisecond)
			counter++
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}

type myMutex struct {
	ch chan struct{}
}

func makeMutex() *myMutex {
	return &myMutex{
		ch: make(chan struct{}, 1),
	}
}

func (m *myMutex) Lock() {
	m.ch <- struct{}{}
}

func (m *myMutex) Unlock() {
	<-m.ch
}

type Item[T any] struct {
	Value T
	Next  *Item[T]
}

func Reverse[T any](iter *Item[T]) *Item[T] {
	var prev *Item[T]
	for iter != nil {
		iter.Next, prev, iter = prev, iter, iter.Next
	}

	return prev
}

func (i *Item[T]) Print() {
	iter := i
	for iter != nil {
		println(iter.Value)
		iter = iter.Next
	}
}
