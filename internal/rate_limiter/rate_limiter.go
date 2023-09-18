package ratelimiter

import (
	"sync"
	"sync/atomic"
	"time"
)

type limiter struct {
	counter         int64
	burst           int64
	limit           int64
	lastSeenSeconds int64
}

func (l *limiter) allow() (ok bool) {
	v := atomic.AddInt64(&l.counter, 1)
	if v <= l.limit+l.burst {
		ok = true
	} else {
		atomic.AddInt64(&l.counter, -1)
	}

	if v >= l.limit {
		lastSeen := atomic.LoadInt64(&l.lastSeenSeconds)
		now := time.Now().Unix()
		if lastSeen < now && atomic.CompareAndSwapInt64(&l.lastSeenSeconds, lastSeen, now) {
			if lastSeen+1 == now {
				atomic.AddInt64(&l.counter, -l.limit)
			} else {
				atomic.StoreInt64(&l.counter, 0)
			}
		}
	}
	return
}

type limiter2 struct {
	limit    int64
	per      int64
	windows  [2]int64
	lastseen int64
	mu       sync.Mutex
}

func NewLimiter2(limit int64, per time.Duration) *limiter2 {
	return &limiter2{
		limit:    limit,
		per:      int64(per),
		windows:  [2]int64{0, 0},
		lastseen: time.Now().UnixNano(),
		mu:       sync.Mutex{},
	}
}

func (l *limiter2) allow() bool {
	var counter int64

	now := time.Now().UnixNano()
	delta := now - l.lastseen

	curr := delta / l.per & 1
	prev := delta/l.per&1 ^ 1

	l.mu.Lock()
	defer l.mu.Unlock()

	switch {
	case delta >= 2*l.per:
		l.windows[prev] = 0

		fallthrough
	case delta >= l.per:
		counter = 1
		l.windows[curr] = counter
		l.lastseen = now
	default:
		l.windows[curr]++
		counter = l.windows[curr]
		counter += l.windows[prev] * (l.per - delta) / l.per
	}

	if counter > l.limit {
		l.windows[curr]--
	}

	return counter <= l.limit
}

type rateLimiter struct {
	values sync.Map
	burst  int64
	limit  int64
}

func (r *rateLimiter) cleanUp() {
	timestamp := time.Now().Unix() - 60
	r.values.Range(func(key, value interface{}) bool {
		lim, ok := value.(limiter)
		if !ok {
			r.values.Delete(key)
		}
		if lim.lastSeenSeconds < timestamp {
			r.values.Delete(key)
		}
		return true
	})
}

func (r *rateLimiter) allow(id interface{}) bool {
	v, ok := r.values.Load(id)
	if !ok {
		v, _ = r.values.LoadOrStore(id, &limiter{burst: r.burst, limit: r.limit})
	}

	lim, ok := v.(*limiter)
	if !ok {
		panic("fail rate limit type cast")
	}
	return lim.allow()
}

func newLimiter(limit, burst int64) *rateLimiter {
	l := &rateLimiter{
		burst: burst,
		limit: limit,
	}
	go func() {
		for range time.NewTicker(time.Minute).C {
			l.cleanUp()
		}
	}()

	return l
}
