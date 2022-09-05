# atomiccounter_bench

This is the benchmark for different atomic counter libraries.

See [atomiccounter](https://github.com/chen3feng/atomiccounter) for details.

## Tested Libraries

- https://github.com/chen3feng/atomiccounter
- https://github.com/puzpuzpuz/xsync
- https://github.com/linxGnu/go-adder
- https://github.com/line/garr

## Test Result

The flowing is one of a typical test result, the format is manual adjusted for easy comparasion.

Test environment is a MacBookPro with MacOS 12.5.1 and Apple M1 Pro chip.

```console
% go test -bench .
goos: darwin
goarch: arm64
pkg: atomiccounter_bench

BenchmarkAdd_NonAtomic-10               49337793                22.02 ns/op
BenchmarkAdd_Atomic-10                    206678                 6854 ns/op
BenchmarkAdd_AtomicCounter-10           14658782                82.22 ns/op
BenchmarkAdd_XsyncCounter-10             9599529                144.6 ns/op
BenchmarkAdd_GoAdder-10                   825858                 1339 ns/op
BenchmarkAdd_GarrAdder-10                 915090                 1305 ns/op

BenchmarkRead_NonAtomic-10             263460258                4.087 ns/op
BenchmarkRead_Atomic-10                172530186                6.945 ns/op
BenchmarkRead_AtomicCounter-10           2793618                425.2 ns/op
BenchmarkRead_XSyncCounter-10            2396407                489.6 ns/op
BenchmarkRead_GoAdder-10                32101244                36.02 ns/op
BenchmarkRead_GarrAdder-10              29420326                35.40 ns/op

PASS
ok      atomiccounter_bench     17.824s
```

As you can see, `atomiccounter` is the fastest for writing (which is its major propose).
