# Serialize and Deserialize Benchmark

``` sh
go test -benchmem -run=^-coverprofile=/tmp/vscode-godtJXfF/go-code-cover -bench . mykitex
goos: linux
goarch: amd64
pkg: mykitex
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
BenchmarkProtobufSerialize-16            1489050               818.2 ns/op           224 B/op          1 allocs/op
BenchmarkProtobufDeserialize-16           958778              1189 ns/op             640 B/op         19 allocs/op
BenchmarkPrutalSerialize-16              1999399               580.9 ns/op           504 B/op          6 allocs/op
BenchmarkPrutalDeserialize-16            1474336               810.7 ns/op           597 B/op          1 allocs/op
BenchmarkMsgpSerialize-16                1000000              2194 ns/op            1272 B/op          0 allocs/op
BenchmarkMsgpDeserialize-16              1544336               741.3 ns/op           128 B/op         18 allocs/op
BenchmarkFrugalSerialize-16              3113959               368.1 ns/op           352 B/op          1 allocs/op
BenchmarkFrugalDeserialize-16            1637936               719.9 ns/op           597 B/op          1 allocs/op
PASS
ok      mykitex 14.657s
```

*Frugal can only run on `amd64` architecture