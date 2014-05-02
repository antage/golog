# golog

The package is logging library for Go.

## Features

* Multiple backends: io.Writer, os.Stderr/os.Stdout, log/syslog.
* Multiple chained formatters: Noop, Color, Newline, etc.
* Each backend can have own formatter and severity mask.

## How to install

```
go get github.com/antage/golog
```

## Usage

```go
package main

import (
    "github.com/antage/golog"
)

var log golog.Logger

func init() {
    log = golog.NewLogger()

    log.AddBackend(
        golog.ERROR,
        golog.DefaultConsoleFormatter(),
        golog.NewConsoleBackend())

    syslog, err := golog.NewSyslogBackend(golog.SYSLOG_LOCAL0, "exampled")
    if err != nil {
        panic(err.Error())
    }

    log.AddBackend(
        golog.ERROR | golog.INFO,
        golog.NoopFormatter{},
        syslog)
}

func main() {
    x := 1
    log.Debugf("x = %d", x)
    log.Infof("x has value %d", x)
    log.Errorf("some error")
    log.Fatalf("some error and os.Exit(1)")
}
```

## Documentation

http://godoc.org/github.com/antage/golog

## License

See LICENSE file.
