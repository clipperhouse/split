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

Some initial benchmarks:

`split.String` (this package)

```
1185 ns/op	    404.28 MB/s	       0 B/op	       0 allocs/op
```

`strings.Split` (standard library)

```
1267 ns/op	    378.07 MB/s	    1280 B/op	       1 allocs/op
```

Not ready for production yet! More testing and API consideration to come.

PR's are welcome, perhaps you'd like to [implement a range iterator](https://tip.golang.org/doc/go1.23).
