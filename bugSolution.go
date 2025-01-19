```go
package main

import "fmt"

func main() {
    var m map[string]int
    value, ok := m["key"]
    if ok {
        fmt.Println("Value found:", value)
    } else {
        fmt.Println("Key not found") // Handle the case where the key doesn't exist
    }
}
```