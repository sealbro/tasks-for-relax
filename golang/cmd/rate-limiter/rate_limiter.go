package main

import (
	"sync"
	"time"
)

type userCounter struct {
	count              int
	start              time.Time
	requestsPerUserKey int
	timeFrame          time.Duration
}

// FixedWindowWithUserDefinedStartRateLimiter Fixed window with user-defined start rate limiter algorithm
type FixedWindowWithUserDefinedStartRateLimiter struct {
	l                      sync.Mutex
	counter                map[string]*userCounter
	baseRequestsPerUserKey int
	baseTimeFrame          time.Duration
}

func New(requestsPerUserKey int, timeFrame time.Duration) *FixedWindowWithUserDefinedStartRateLimiter {
	return &FixedWindowWithUserDefinedStartRateLimiter{
		baseRequestsPerUserKey: requestsPerUserKey,
		baseTimeFrame:          timeFrame,
		counter:                make(map[string]*userCounter),
	}
}

// Set custom setup for specific userKey
func (r *FixedWindowWithUserDefinedStartRateLimiter) Set(userKey string, requestsPerUserKey int, timeFrame time.Duration) {
	r.l.Lock()
	defer r.l.Unlock()

	if v, ok := r.counter[userKey]; ok {
		v.requestsPerUserKey = requestsPerUserKey
		v.timeFrame = timeFrame
	} else {
		r.counter[userKey] = &userCounter{
			count:              0,
			start:              time.Time{},
			requestsPerUserKey: requestsPerUserKey,
			timeFrame:          timeFrame,
		}
	}
}

// Allow checks if the rate limiter allows the request
func (r *FixedWindowWithUserDefinedStartRateLimiter) Allow(userKey string) bool {
	r.l.Lock()
	defer r.l.Unlock()

	if v, ok := r.counter[userKey]; ok {
		if time.Since(v.start) > v.timeFrame {
			v.count = 1
			v.start = time.Now()
			return true
		}
		if v.count < v.requestsPerUserKey {
			v.count++
			return true
		}
		return false
	}

	r.counter[userKey] = &userCounter{
		count:              1,
		start:              time.Now(),
		requestsPerUserKey: r.baseRequestsPerUserKey,
		timeFrame:          r.baseTimeFrame,
	}
	return true
}

// Reset resets the rate limiter for a specific userKey to the base configuration
func (r *FixedWindowWithUserDefinedStartRateLimiter) Reset(userKey string) {
	r.l.Lock()
	defer r.l.Unlock()
	if v, ok := r.counter[userKey]; ok {
		v.count = 0
		v.start = time.Time{}
		v.requestsPerUserKey = r.baseRequestsPerUserKey
		v.timeFrame = r.baseTimeFrame
	}
}

// ResetAll resets all rate limiters
func (r *FixedWindowWithUserDefinedStartRateLimiter) ResetAll() {
	r.l.Lock()
	defer r.l.Unlock()
	r.counter = make(map[string]*userCounter)
}
