package main

import (
	"testing"
	"time"
)

func TestFixedWindowWithUserDefinedStartRateLimiter_Allow_Same(t *testing.T) {
	limiter := New(10, time.Millisecond*100)
	allowed := map[string]int{
		"user1": 0,
		"user2": 0,
		"user3": 0,
	}

	for user, _ := range allowed {
		for i := 0; i < 30; i++ {
			if limiter.Allow(user) {
				allowed[user]++
			}
		}
	}

	for user, _ := range allowed {
		if allowed[user] != 10 {
			t.Errorf("Expected 10, got %d", allowed[user])
		}
	}

	time.Sleep(time.Millisecond * 100)

	for user, _ := range allowed {
		for i := 0; i < 30; i++ {
			if limiter.Allow(user) {
				allowed[user]++
			}
		}
	}

	for user, _ := range allowed {
		if allowed[user] != 20 {
			t.Errorf("Expected 20, got %d", allowed[user])
		}
	}
}

func TestFixedWindowWithUserDefinedStartRateLimiter_Allow_Different(t *testing.T) {
	limiter := New(10, time.Millisecond*100)
	allowed := map[string]int{
		"user1": 0,
		"user2": 0,
		"user3": 0,
	}

	for user, _ := range allowed {
		for i := 0; i < 30; i++ {
			if limiter.Allow(user) {
				allowed[user]++
			}
		}
	}
	for user, _ := range allowed {
		if allowed[user] != 10 {
			t.Errorf("Expected 10, got %d", allowed[user])
		}
	}

	limiter.Set("user1", 20, time.Millisecond*50)

	time.Sleep(time.Millisecond * 100)

	for user, _ := range allowed {
		for i := 0; i < 30; i++ {
			if limiter.Allow(user) {
				allowed[user]++
			}
		}
	}

	for user, _ := range allowed {
		if user == "user1" {
			if allowed[user] != 30 {
				t.Errorf("Expected 40, got %d", allowed[user])
			}
		} else {
			if allowed[user] != 20 {
				t.Errorf("Expected 20, got %d", allowed[user])
			}
		}
	}
}

func TestFixedWindowWithUserDefinedStartRateLimiter_Allow_Different2(t *testing.T) {
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
	for user, _ := range allowed {
		if user == "user1" {
			if allowed[user] != 20 {
				t.Errorf("Expected 20, got %d", allowed[user])
			}
		} else {
			if allowed[user] != 10 {
				t.Errorf("Expected 10, got %d", allowed[user])
			}
		}
	}

	time.Sleep(time.Millisecond * 100)

	for user, _ := range allowed {
		for i := 0; i < 30; i++ {
			if limiter.Allow(user) {
				allowed[user]++
			}
		}
	}

	for user, _ := range allowed {
		if user == "user1" {
			if allowed[user] != 40 {
				t.Errorf("Expected 40, got %d", allowed[user])
			}
		} else {
			if allowed[user] != 20 {
				t.Errorf("Expected 20, got %d", allowed[user])
			}
		}
	}
}

func TestFixedWindowWithUserDefinedStartRateLimiter_Reset_UserKey(t *testing.T) {
	limiter := New(10, time.Millisecond*100)
	allowed := map[string]int{
		"user1": 0,
		"user2": 0,
		"user3": 0,
	}

	for user, _ := range allowed {
		for i := 0; i < 30; i++ {
			if limiter.Allow(user) {
				allowed[user]++
			}
		}
	}
	limiter.Reset("user1")
	for user, _ := range allowed {
		for i := 0; i < 30; i++ {
			if limiter.Allow(user) {
				allowed[user]++
			}
		}
	}

	for user, _ := range allowed {
		if user == "user1" {
			if allowed[user] != 20 {
				t.Errorf("Expected 10, got %d", allowed[user])
			}
		} else {
			if allowed[user] != 10 {
				t.Errorf("Expected 20, got %d", allowed[user])
			}
		}
	}
}

func TestFixedWindowWithUserDefinedStartRateLimiter_Reset_All(t *testing.T) {
	limiter := New(10, time.Millisecond*100)
	allowed := map[string]int{
		"user1": 0,
		"user2": 0,
		"user3": 0,
	}

	for user, _ := range allowed {
		for i := 0; i < 30; i++ {
			if limiter.Allow(user) {
				allowed[user]++
			}
		}
	}

	limiter.ResetAll()

	for user, _ := range allowed {
		for i := 0; i < 30; i++ {
			if limiter.Allow(user) {
				allowed[user]++
			}
		}
	}

	for user, _ := range allowed {
		if allowed[user] != 20 {
			t.Errorf("Expected 10, got %d", allowed[user])
		}
	}
}
