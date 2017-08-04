package randpool

import (
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

func (p Pool) Next() *rand.Rand {
	i := atomic.AddInt32(&p.iter, 1)
	if i > p.capacity {
		atomic.StoreInt32(&p.iter, 0)
		i = 0
	}
	return p.rands[i]
}
