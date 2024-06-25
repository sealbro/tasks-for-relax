package main

import (
	"fmt"
	"time"
)

func main() {
	limiter := New(10, time.Millisecond*100)
	allowed := map[string]int{
		"user1": 0,
		"user2": 0,
		"user3": 0,
	}
	limiter.Set("user1", 20, time.Millisecond*50)

	for user, _ := range allowed {
		for i := 0; i < 30; i++ {
			if limiter.Allow(user) {
				allowed[user]++
			}
		}
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println(allowed)

	for user, _ := range allowed {
		for i := 0; i < 30; i++ {
			if limiter.Allow(user) {
				allowed[user]++
			}
		}
	}
	fmt.Println(allowed)
}
