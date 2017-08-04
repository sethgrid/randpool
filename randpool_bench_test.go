package randpool

import "testing"
import "time"
import "math/rand"
import "github.com/NebulousLabs/fastrand"

var result int32

func BenchmarkRandPool(b *testing.B) {
	var i int32
	b.StopTimer()
	pool := New(1000, time.Now().UnixNano)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		i = pool.Next().Int31n(10000)
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
