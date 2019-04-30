## LogAdapter

This repo is a adapter to the package [nuveo/log](github.com/nuveo/log) that uses Golang's standard `log` instead of `fmt` package to write into stdout. This is useful for testing if a application logged something by setting the log to a buffer.


## [Testing logs with this adapter](https://stackoverflow.com/questions/44119951/how-to-check-a-log-output-in-go-test)


```go
package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
    "testing"
    l "log"

    "github.com/nuveo/log"
    la "github.com/arxdsilva/logAdapter"
)

func readByte() {
     // force an error
    err := io.EOF
    if err != nil {
        log.Println("Couldn't read first byte")
    }
}

func TestReadByte(t *testing.T) {
    // remove std adapter 
    // this prevents double logging
    log.RemoveAdapter("stdout")
    log.AddAdapter("adapter", log.AdapterPod{Adapter:la.LogAdapter, Config: nil})
    var buf bytes.Buffer
    l.SetOutput(&buf)
    defer func() {
        l.SetOutput(os.Stderr)
    }()
    readByte()
    // do something with the logs
    t.Log(buf.String())
}
```
