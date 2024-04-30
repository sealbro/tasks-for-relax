package benchmarks

import (
	"sync"
	"sync/atomic"
)

const max = 10000

func UseAtomic() {
	var init int64 = 0
	var count *int64
	count = &init
	wg := sync.WaitGroup{}
	wg.Add(max)

	for i := 0; i < max; i++ {
		go func() {
			atomic.AddInt64(count, 1)

			wg.Done()
		}()
	}

	wg.Wait()

	if *count != max {
		panic(*count)
	}
}

func UseLock() {
	var count int64 = 0
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(max)

	for i := 0; i < max; i++ {
		go func() {
			mutex.Lock()
			count++
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	if count != max {
		panic(count)
	}
}

func UseMutexChannel() {
	var count int64 = 0
	mutex := MakeMyMutex()
	wg := sync.WaitGroup{}
	wg.Add(max)

	for i := 0; i < max; i++ {
		go func() {
			mutex.Lock()
			count++
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	mutex.Close()

	if count != max {
		panic(count)
	}
}

func UseChannel() {
	ch := make(chan int64)
	var count int64 = 0
	wg := sync.WaitGroup{}
	wg.Add(max)

	go func() {
		for i := range ch {
			count += i
			wg.Done()
		}
	}()

	for i := 0; i < max; i++ {
		go func() {
			ch <- int64(1)
		}()
	}

	wg.Wait()
	close(ch)

	if count != max {
		panic(count)
	}
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
