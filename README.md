# buoy

> tiny, relational database


```go
package main

import (
    "fmt"
    "github.com/seadreamcode/buoy"
)

func main() {
    db, err := buoy.Connect()
    defer db.Close()

    db.Exec(`CREATE DATABASE test`)
}
```