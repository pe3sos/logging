## Logging

A simple leveled logging library with coloured output.

[![Travis Status for zippunov/logging](https://travis-ci.org/zippunov/logging.svg?branch=master&label=linux+build)](https://travis-ci.org/zippunov/logging)
[![godoc for zippunov/logging](https://godoc.org/github.com/nathany/looper?status.svg)](http://godoc.org/github.com/zippunov/logging)

---

Log levels:
- `INFO` (grey)
- `INFO` (blue)
- `WARNING` (pink)
- `ERROR` (red)
- `FATAL` (red)

Formatters:

- `Default`
- `Coloured`
- `Compose`
- `Date`
- `Level`
- `Location`
- `Time`
- `Timestamp`

Example usage. Create a new package `log` in your app such that:

```go
package log

import (
	"os"

	"github.com/zippunov/logging"
	"github.com/zippunov/logging/def"
	"github.com/zippunov/logging/formatter"
)

var f = formatter.Compose(
	formatter.Coloured(),
	formatter.Level(),
	formatter.Timestamp(),
	formatter.Location(5),
)

var (
	logger  = logging.New(os.Stdout, os.Stdout, f)
	DEBUG   = logger.DEBUG
	INFO    = logger.INFO
	WARNING = logger.WARNING
	ERROR   = logger.ERROR
	FATAL   = logger.FATAL
)

func SetLevel(lvl def.Level) {
	logger.SetMaxLevel(lvl)
}
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
