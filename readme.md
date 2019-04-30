## LogAdapter

This repo is a adapter to the package `nuveo/log` that uses Golang's standard `log` instead of `fmt` package to write into stdout. This is useful for testing if a application logged something by setting the log to a buffer.


## [Testing logs with this adapter](https://stackoverflow.com/questions/44119951/how-to-check-a-log-output-in-go-test)


```go
package main

import (
    "bytes"
    "fmt"
    "io"
    // using `arxdsilva` and not `nuveo` due to a needed typo fix 
    "github.com/arxdsilva/log"
    "github.com/arxdsilva/logAdapter"
    "os"
    "testing"
)

func readByte() {
     // force an error
    err := io.EOF
    if err != nil {
        log.Println("Couldn't read first byte")
    }
}

func TestReadByte(t *testing.T) {
    log.RemoveAdapter("stdout")
    log.AddAdapter("adapter", log.AdapterPod{Adapter:la.LogAdapter, Config: nil})
    var buf bytes.Buffer
    log.SetOutput(&buf)
    defer func() {
        log.SetOutput(os.Stderr)
    }()
    readByte()
    t.Log(buf.String())
}
```
