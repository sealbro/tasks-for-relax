package benchmarks

import "sync"

const max = 10000

func UseLock() {
	count := 0
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
}

func UseChannel() {
	ch := make(chan struct{})
	count := 0
	wg := sync.WaitGroup{}
	wg.Add(max)

	go func() {
		for range ch {
			count++
			wg.Done()
		}
	}()

	for i := 0; i < max; i++ {
		go func() {
			ch <- struct{}{}
		}()
	}

	wg.Wait()
	close(ch)
}
