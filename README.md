# Go concurrency experiments

Test scenarios:

1) SPMC (between 1-5 readers) reading/writing to a single variable 

    1) no synchronisation
    2) mutex
    3) rwmutex
    4) atomic

```
goos: linux
goarch: amd64
pkg: concurrencybenchmark
cpu: AMD Ryzen 7 7700X 8-Core Processor
BenchmarkGlobalVariable/readers:_1-16   822691628                1.335 ns/op           0 B/op          0 allocs/op
BenchmarkGlobalVariable/readers:_2-16   774267472                1.397 ns/op           0 B/op          0 allocs/op
BenchmarkGlobalVariable/readers:_3-16   963245335                1.208 ns/op           0 B/op          0 allocs/op
BenchmarkGlobalVariable/readers:_4-16   876295995                1.348 ns/op           0 B/op          0 allocs/op
BenchmarkGlobalVariable/readers:_5-16   661642039                1.817 ns/op           0 B/op          0 allocs/op
BenchmarkVariableMutexReadLock/readers:_1-16            169302352                7.207 ns/op           0 B/op          0 allocs/op
BenchmarkVariableMutexReadLock/readers:_2-16            40004056                32.89 ns/op            0 B/op          0 allocs/op
BenchmarkVariableMutexReadLock/readers:_3-16            17158455                61.20 ns/op            0 B/op          0 allocs/op
BenchmarkVariableMutexReadLock/readers:_4-16            12648674                92.48 ns/op            0 B/op          0 allocs/op
BenchmarkVariableMutexReadLock/readers:_5-16             9363747               125.4 ns/op             0 B/op          0 allocs/op
BenchmarkVariableRWMutexReadLock/readers:_1-16           5853928               198.7 ns/op             0 B/op          0 allocs/op
BenchmarkVariableRWMutexReadLock/readers:_2-16           6826298               174.5 ns/op             0 B/op          0 allocs/op
BenchmarkVariableRWMutexReadLock/readers:_3-16           6085744               190.8 ns/op             0 B/op          0 allocs/op
BenchmarkVariableRWMutexReadLock/readers:_4-16           5570672               212.7 ns/op             0 B/op          0 allocs/op
BenchmarkVariableRWMutexReadLock/readers:_5-16           5220693               224.7 ns/op             0 B/op          0 allocs/op
BenchmarkVariableAtomic/readers:_1-16                   1000000000               1.357 ns/op           0 B/op          0 allocs/op
BenchmarkVariableAtomic/readers:_2-16                   997141653                1.175 ns/op           0 B/op          0 allocs/op
BenchmarkVariableAtomic/readers:_3-16                   1000000000               0.9973 ns/op          0 B/op          0 allocs/op
BenchmarkVariableAtomic/readers:_4-16                   1000000000               0.9413 ns/op          0 B/op          0 allocs/op
BenchmarkVariableAtomic/readers:_5-16                   1000000000               1.236 ns/op           0 B/op          0 allocs/op
BenchmarkRWMutexRealistic/readers:_1-16                 323110401                3.826 ns/op           0 B/op          0 allocs/op
BenchmarkRWMutexRealistic/readers:_2-16                 61201333                20.53 ns/op            0 B/op          0 allocs/op
BenchmarkRWMutexRealistic/readers:_3-16                 84495164                18.03 ns/op            0 B/op          0 allocs/op
BenchmarkRWMutexRealistic/readers:_4-16                 59269488                23.86 ns/op            0 B/op          0 allocs/op
BenchmarkRWMutexRealistic/readers:_5-16                 41407112                29.42 ns/op            0 B/op          0 allocs/op
```

2) MPMC reading/writing to a single variable 

    1) no synchronisation
    2) mutex
    3) rwmutex
    4) atomic

3) low writes/heavy reads

    ...

4) Complex objects - maps/slices/channels

    1) A mutuex protected map vs sync.Map
    2) A channel vs mutex protected deque
    3) A channel vs a ringbuffer

