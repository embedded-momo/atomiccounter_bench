package atomiccounterbench

import (
	"sync/atomic"
	"testing"

	"github.com/chen3feng/atomiccounter"
	ga "github.com/linxGnu/go-adder"
	"github.com/puzpuzpuz/xsync"
	garr "go.linecorp.com/garr/adder"
)

//go:norace
//go:noinline
func add(count *int64, n int) {
	for i := 0; i < n; i++ {
		*count++
	}
}

func atomicAdd(count *int64, n int) {
	for i := 0; i < n; i++ {
		atomic.AddInt64(count, 1)
	}
}

const (
	batchSize = 100
)

func BenchmarkAdd_NonAtomic(b *testing.B) {
	b.SetParallelism(100)
	count := int64(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			add(&count, batchSize)
		}
	})
}

func BenchmarkAdd_Atomic(b *testing.B) {
	b.SetParallelism(100)
	count := int64(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomicAdd(&count, batchSize)
		}
	})
}

func BenchmarkAdd_AtomicCounter(b *testing.B) {
	b.SetParallelism(100)
	counter := atomiccounter.MakeInt64()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < batchSize; i++ {
				counter.Add(1)
			}
		}
	})
}

func BenchmarkAdd_XsyncCounter(b *testing.B) {
	b.SetParallelism(100)
	counter := xsync.Counter{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < batchSize; i++ {
				counter.Add(1)
			}
		}
	})
}

func BenchmarkAdd_GoAdder(b *testing.B) {
	b.SetParallelism(100)
	counter := ga.NewLongAdder(ga.JDKAdderType)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < batchSize; i++ {
				counter.Add(1)
			}
		}
	})
}

func BenchmarkAdd_GarrAdder(b *testing.B) {
	b.SetParallelism(100)
	counter := garr.NewLongAdder(garr.JDKAdderType)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < batchSize; i++ {
				counter.Add(1)
			}
		}
	})
}

func BenchmarkRead_NonAtomic(b *testing.B) {
	count := int64(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < batchSize; i++ {
				_ = count
			}
		}
	})
}

func BenchmarkRead_Atomic(b *testing.B) {
	count := int64(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < batchSize; i++ {
				_ = atomic.LoadInt64(&count)
			}
		}
	})
}

func BenchmarkRead_AtomicCounter(b *testing.B) {
	counter := atomiccounter.MakeInt64()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < batchSize; i++ {
				_ = counter.Read()
			}
		}
	})
}

func BenchmarkRead_XSyncCounter(b *testing.B) {
	counter := xsync.Counter{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < batchSize; i++ {
				_ = counter.Value()
			}
		}
	})
}

func BenchmarkRead_GoAdder(b *testing.B) {
	counter := ga.NewLongAdder(ga.JDKAdderType)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < batchSize; i++ {
				_ = counter.Sum()
			}
		}
	})
}

func BenchmarkRead_GarrAdder(b *testing.B) {
	counter := garr.NewLongAdder(garr.JDKAdderType)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < batchSize; i++ {
				_ = counter.Sum()
			}
		}
	})
}
