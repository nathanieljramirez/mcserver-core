## Installing

```bash
go get mcserverkit.github.io/core@latest
```

## Usage

```go
package main

import (
  "mcserverkit.github.io/core"
  "fmt"
)

func main() {
  err := core.Install("1.21.1")
  if err != nil {
    fmt.Println("Error installing 1.21.1:", err)
    return
  }

  err = core.Create("MyServer")
  if err != nil {
    fmt.Println("Error creating MyServer:", err)
    return
  }

  err = core.Start("MyServer", "4G")
  if err != nil {
    fmt.Println("Error starting MyServer:", err)
    return
  }
}
```
