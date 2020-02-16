## Logging

A simple leveled logging library with coloured output.

[![Travis Status for zippunov/logging](https://travis-ci.org/zippunov/logging.svg?branch=master&label=linux+build)](https://travis-ci.org/zippunov/logging)
[![godoc for zippunov/logging](https://godoc.org/github.com/nathany/looper?status.svg)](http://godoc.org/github.com/zippunov/logging)

---

Log levels:

- `INFO` (blue)
- `WARNING` (pink)
- `ERROR` (red)
- `FATAL` (red)

Formatters:

- `DefaultFormatter`
- `ColouredFormatter`

Example usage. Create a new package `log` in your app such that:

```go
package log

import (
	"log"
	"github.com/zippunov/logging"
)

var (
	logger = logging.New(nil, nil, new(logging.ColouredFormatter), log.Ldate|log.Ltime)

	// INFO ...
	INFO = logger[logging.INFO]
	// WARNING ...
	WARNING = logger[logging.WARNING]
	// ERROR ...
	ERROR = logger[logging.ERROR]
	// FATAL ...
	FATAL = logger[logging.FATAL]
)
```

Then from your app you could do:

```go
package main

import (
	"github.com/yourusername/yourapp/log"
)

func main() {
	log.INFO.Print("log message")
}
```
