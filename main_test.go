package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

const (
	minReaders = 1
	maxReaders = 5
)

// Benchmark 1: No synchronization (baseline)
// Go detects race conditions here, but it's only a baseline test.
//
//go:nocheckrace
func BenchmarkGlobalVariable(b *testing.B) {
	for readers := minReaders; readers <= maxReaders; readers++ {
		b.Run(fmt.Sprintf("readers: %d", readers), func(b *testing.B) {
			var noSyncValue int64
			var wg sync.WaitGroup

			// Start the single writer
			done := make(chan struct{}) // To signal when to stop
			go func() {
				for i := 0; ; i++ {
					select {
					case <-done:
						return
					default:
						noSyncValue = int64(i)
					}
				}
			}()

			b.ResetTimer()

			// Start `readers` goroutines
			for n := 1; n <= readers; n++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					var localValue int64

					for i := 0; i < b.N; i++ {
						localValue = noSyncValue
					}

					runtime.KeepAlive(localValue)
				}()
			}

			wg.Wait()
			close(done)
		})
	}
}

// Benchmark 2: Mutex-protected read and write access
func BenchmarkVariableMutexReadLock(b *testing.B) {
	for readers := minReaders; readers <= maxReaders; readers++ {
		b.Run(fmt.Sprintf("readers: %d", readers), func(b *testing.B) {
			var muProtectedValue int64
			var mu sync.Mutex
			var wg sync.WaitGroup

			// Launch one writer goroutine
			done := make(chan struct{}) // To signal when to stop

			go func() {
				for i := 0; ; i++ {
					select {
					case <-done:
						return
					default:
						mu.Lock()
						muProtectedValue = int64(i)
						mu.Unlock()
					}
				}
			}()

			b.ResetTimer()

			// Reader in main goroutine
			for n := 1; n <= readers; n++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for x := 0; x < b.N; x++ {
						var localValue int64

						mu.Lock()
						localValue = muProtectedValue
						mu.Unlock()

						runtime.KeepAlive(localValue)
					}
				}()
			}
			wg.Wait()
			close(done)
		})
	}
}

// Benchmark 2: RWMutex-protected allowing concurrent reads
func BenchmarkVariableRWMutexReadLock(b *testing.B) {
	for readers := minReaders; readers <= maxReaders; readers++ {
		b.Run(fmt.Sprintf("readers: %d", readers), func(b *testing.B) {
			var muProtectedValue int64
			var mu sync.RWMutex
			var wg sync.WaitGroup

			// Launch one writer goroutine
			done := make(chan struct{}) // To signal when to stop

			go func() {
				for i := 0; ; i++ {
					select {
					case <-done:
						return
					default:
						mu.Lock()
						muProtectedValue = int64(i)
						mu.Unlock()
					}
				}
			}()

			b.ResetTimer()

			// Reader in main goroutine
			for n := 1; n <= readers; n++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for x := 0; x < b.N; x++ {
						var localValue int64

						mu.RLock()
						localValue = muProtectedValue
						mu.RUnlock()

						runtime.KeepAlive(localValue)
					}
				}()
			}
			wg.Wait()
			close(done)
		})
	}
}

// Benchmark 4: Atomic access
func BenchmarkVariableAtomic(b *testing.B) {
	for readers := minReaders; readers <= maxReaders; readers++ {
		b.Run(fmt.Sprintf("readers: %d", readers), func(b *testing.B) {
			var atomicValue atomic.Int64
			var wg sync.WaitGroup

			done := make(chan struct{})
			b.ResetTimer()
			go func() {
				for i := 0; ; i++ {
					select {
					case <-done:
						return
					default:
						atomicValue.Store(int64(i))
					}
				}
			}()

			for n := 1; n <= readers; n++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					var localValue int64

					for i := 0; i < b.N; i++ {
						localValue = atomicValue.Load()
					}

					runtime.KeepAlive(localValue)
				}()
			}

			wg.Wait()
			close(done)
		})
	}
}

// Benchmark 5: RWMutex with realistic write frequency
func BenchmarkRWMutexRealistic(b *testing.B) {
	for readers := minReaders; readers <= maxReaders; readers++ {
		b.Run(fmt.Sprintf("readers: %d", readers), func(b *testing.B) {
			var rwProtectedValue int64
			var rwMu sync.RWMutex
			var wg sync.WaitGroup

			// Writer that only writes occasionally (1% of the time)
			go func() {
				for i := 0; ; i++ {
					if i%100 == 0 { // Write only 1% of the time
						rwMu.Lock()
						rwProtectedValue = int64(i)
						rwMu.Unlock()
					}
					runtime.Gosched() // Give other goroutines a chance
				}
			}()

			b.ResetTimer()

			for n := 1; n <= readers; n++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					var localValue int64

					for i := 0; i < b.N; i++ {
						rwMu.RLock()
						localValue = rwProtectedValue
						rwMu.RUnlock()
					}

					runtime.KeepAlive(localValue)
				}()
			}

			wg.Wait()
		})
	}
}
