```go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["key"]) // This will not panic, but it will print 0.
}
```