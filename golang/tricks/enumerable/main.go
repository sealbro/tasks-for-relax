package main

import (
	"sync"
	"time"
)

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
		close(ch)
	}()

	return ch
}

func main() {
	timeStart := time.Now()

	ch1 := worker()
	ch2 := worker()
	closed1 := false
	closed2 := false
	for {
		select {
		case _, cl1 := <-ch1:
			closed1 = closed1 || cl1
		case _, cl2 := <-ch2:
			closed2 = closed2 || cl2
		default:
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if closed1 && closed2 {
			break
		}
	}

	//_, _ = <-worker(), <-worker()

	println(int(time.Since(timeStart).Seconds())) // что выведет - 3 или 6?
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

type A struct {
}

func (A) Show() {
	println("A")
}

type B struct {
	A
}

func (B) Show() {
	println("B")
}

type MyMutex struct {
	ch chan struct{}
}

func MakeMyMutex() *MyMutex {
	return &MyMutex{ch: make(chan struct{}, 1)}
}

func (m *MyMutex) Close() {
	close(m.ch)
}

func (m *MyMutex) Lock() {
	m.ch <- struct{}{}

}

func (m *MyMutex) Unlock() {
	<-m.ch
}

func main2() {
	m := MakeMyMutex()
	defer m.Close()

	ln := 1000
	counter := 0
	group := sync.WaitGroup{}
	group.Add(ln)
	for i := 0; i < ln; i++ {
		go func() {
			defer group.Done()
			m.Lock()
			counter++
			defer m.Unlock()
		}()
	}
	group.Wait()

	println(counter)
}

func pow(chs <-chan int) <-chan int {
	pow := make(chan int)

	go func() {
		for i := range chs {
			pow <- i * i
		}
		close(pow)
	}()

	return pow
}

func join(chs ...<-chan int) <-chan int {
	joined := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(chs))
	go func() {
		for i := range chs {
			go func(ch <-chan int) {
				for j := range ch {
					joined <- j
				}
				wg.Done()
			}(chs[i])
		}
		wg.Wait()
		close(joined)
	}()

	return joined
}

func generate(count int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}
