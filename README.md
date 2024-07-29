A more efficient splitter for bytes and strings, with a focus on zero allocation, for Go. Use this where you would might use `bytes.Split` or `strings.Split`.

### Usage

See [pkg.go.dev](https://pkg.go.dev/github.com/clipperhouse/split)

```bash
go get github.com/clipperhouse/split
```

```go
import "github.com/clipperhouse/split"

text := "Hello, 世界. Nice dog! 👍🐶"
sep := " "

split := split.String(text, sep)

for split.Next() {
    fmt.Println(split.Value())
}
```

### Performance

Some initial benchmarks:

`split.String` (this package)

```
1185 ns/op	    404.28 MB/s	       0 B/op	       0 allocs/op
```

`strings.Split` (standard library)

```
1267 ns/op	    378.07 MB/s	    1280 B/op	       1 allocs/op
```

Overall, this package is a _little_ faster, but more importantly, notice the difference in allocations. If you're on a hot path, this might add up, and reducing GC might help your app to scale.

### Why you might use this

The standard library collects all the splits at once into an array, and allocates to do so (this is true in other languages as well).

This package lazily iterates over each split as needed, and avoids that allocation. Think of it as streaming instead of batching.

If you do not actually need the array, but only need to iterate over the splits, this package may be useful.

### Data types

This packages handles `string` and `[]byte` (and named types based on them). If you have an `io.Reader`, we suggest [`bufio.Scanner`](https://pkg.go.dev/bufio) from the standard library.

### Testing

We work to ensure that `split.Bytes` and `split.String` offer an identical API and results as their standard library counterparts, `bytes.Split` and `strings.Split`. Have a look at the [tests](https://github.com/clipperhouse/split/blob/main/tests_test.go) to verify that this is true.

[![Test](https://github.com/clipperhouse/split/actions/workflows/gotest.yml/badge.svg)](https://github.com/clipperhouse/split/actions/workflows/gotest.yml)

### Status

We've published a v0.1.0. Try it and leave feedback.

PR's are welcome, perhaps you'd like to [implement a range iterator](https://tip.golang.org/doc/go1.23).
