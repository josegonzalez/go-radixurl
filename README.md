go-radixurl
============

Connect to a redis instance using a `REDIS_URL`.

Usage
=====

It uses [Radix Redis][radix] under the hood:

```go
import "github.com/josegonzalez/go-radixurl"

// Connect using os.Getenv("REDIS_URL").
c, err := radixurl.Connect()

// Alternatively, connect using a custom Database URL.
c, err := radixurl.ConnectToURL("redis://...")
```

In both cases you will get the result values of calling
`redis.DialTimeout(...)`, that is, an instance of `redis.Client` and an
error.

[radix]: https://github.com/fzzy/radix

Installation
============

Install it using the "go get" command:

    go get github.com/josegonzalez/go-radixurl
