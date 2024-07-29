A more efficient splitter for bytes and strings, with a focus on zero allocation, for Go.

Use this where you would might use `bytes.Split` or `strings.Split`.

```
go get https://github.com/clipperhouse/split
```

```go
import "github.com/clipperhouse/split"
```

```go
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

Overall, this package is a bit faster, but perhaps more importantly, notice the difference in allocations. `strings.Split` seems to have some sort of multiplying effect, as this benchmark uses a 400B string.

### Testing

We work to ensure that `split.Bytes` and `split.String` offer an identical API and results as their standard library counterparts, `bytes.Split` and `strings.Split`. Have a look at the tests to verify that this is true.

[![Test](https://github.com/clipperhouse/split/actions/workflows/gotest.yml/badge.svg)](https://github.com/clipperhouse/split/actions/workflows/gotest.yml)

### Status

Not ready for production yet! More testing and API consideration to come.

PR's are welcome, perhaps you'd like to [implement a range iterator](https://tip.golang.org/doc/go1.23).
