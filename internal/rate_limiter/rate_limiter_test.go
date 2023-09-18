package ratelimiter

import (
	"testing"
	"time"

	"golang.org/x/time/rate"

	"github.com/stretchr/testify/assert"
)

func Test_ratelimiter_default(t *testing.T) {
	lim := newLimiter(10, 20)
	valid := false
	for i := 0; i < 10; i++ {
		valid = lim.allow(1)
	}
	assert.Equal(t, valid, true)
	for i := 0; i < 10; i++ {
		valid = lim.allow(1)
	}
	assert.Equal(t, valid, true)
	for i := 0; i < 30; i++ {
		valid = lim.allow(1)
	}
	assert.Equal(t, valid, false)
	time.Sleep(time.Second)
	assert.Equal(t, lim.allow(1), false)
	time.Sleep(time.Second * 2)
	assert.Equal(t, lim.allow(1), true)
}

func Test_ratelimiter_my(t *testing.T) {
	lim := NewLimiter2(10, time.Second)
	valid := false
	for i := 0; i < 10; i++ {
		valid = lim.allow()
	}
	assert.Equal(t, valid, true)
	for i := 0; i < 10; i++ {
		valid = lim.allow()
	}
	assert.Equal(t, valid, false)
	for i := 0; i < 30; i++ {
		valid = lim.allow()
	}
	assert.Equal(t, valid, false)
	time.Sleep(time.Second)
	assert.Equal(t, lim.allow(), true)
	time.Sleep(time.Second * 2)
	assert.Equal(t, lim.allow(), true)
}

func BenchmarkUs(b *testing.B) {
	lim := limiter{
		burst: 5,
		limit: 10,
	}
	for i := 0; i < b.N; i++ {
		lim.allow()
	}
}

func BenchmarkStock(b *testing.B) {
	lim := rate.NewLimiter(10, 5)
	for i := 0; i < b.N; i++ {
		lim.Allow()
	}
}

func BenchmarkMy(b *testing.B) {
	lim := NewLimiter2(10, time.Second)
	for i := 0; i < b.N; i++ {
		lim.allow()
	}
}
