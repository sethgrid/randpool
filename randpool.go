package randpool

import (
	"math"
	"math/rand"
	"sync/atomic"
)

type Pool struct {
	rands    []*rand.Rand
	iter     int32
	capacity int32
}

func New(size int32, fn func() int64) Pool {
	rands := make([]*rand.Rand, size)
	for i := int32(0); i < size; i++ {
		rands[i] = rand.New(rand.NewSource(fn()))

	}
	return Pool{
		rands:    rands,
		capacity: size,
	}
}

func (p *Pool) Next() *rand.Rand {
	i := atomic.AddInt32(&p.iter, 1)
	if i >= p.capacity {
		atomic.StoreInt32(&p.iter, 0)
		i = 0
	}
	return p.rands[i]
}

var clockID int64

// Clock is a type of lamport clock
type Clock struct {
	ID              int64
	IDWidth         int64
	widthMultiplier int64
	count           int64
}

func NewClock(IDWidth int64) Clock {
	id := atomic.AddInt64(&clockID, 1)
	// if we want automatic id generation, we can remove the passed in IDWidth and use the following
	// IDWidth := int64(math.Log10(float64(id))) + 1
	return Clock{
		ID:              id,
		IDWidth:         IDWidth,
		widthMultiplier: int64(math.Pow(10, float64(IDWidth))),
	}
}

func (c *Clock) Next() int64 {
	// TODO: overflow detection
	return atomic.AddInt64(&c.count, 1)*c.widthMultiplier + c.ID
}
