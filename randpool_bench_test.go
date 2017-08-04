package randpool

import "testing"
import "time"
import "math/rand"
import "github.com/NebulousLabs/fastrand"

var result int32
var res64 int64

func BenchmarkRandPool(b *testing.B) {
	var i int32
	b.StopTimer()
	pool := New(2, time.Now().UnixNano)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		i = pool.Next().Int31n(10000)
	}
	result = i
}

func BenchmarkRandInstancePanic(b *testing.B) {
	b.Skip("un skip to see a panic. yay, rand instance not thread safe.")
	var i int32
	b.StopTimer()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		go func() { r.Int31n(10000) }()
	}
	result = i
}

func BenchmarkRandInstance(b *testing.B) {
	var i int32
	b.StopTimer()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		i = r.Int31n(10000)
	}
	result = i
}

func BenchmarkRandPackage(b *testing.B) {
	var i int32
	for n := 0; n < b.N; n++ {
		i = rand.Int31n(10000)
	}
	result = i
}

func BenchmarkFastRand(b *testing.B) {
	var i int
	for n := 0; n < b.N; n++ {
		i = fastrand.Intn(10000)
	}
	result = int32(i)
}

func BenchmarkClock(b *testing.B) {
	var i int64
	c := NewClock(4)
	for n := 0; n < b.N; n++ {
		i = c.Next()
	}
	res64 = i
}
