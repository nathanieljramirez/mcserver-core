## Installing

```bash
go get mcserverkit.github.io/core@latest
```

## Usage

```go
package main

import (
  "mcserverkit.github.io/core"
)

core.Install("1.21.1")
core.Create("MyServer")
core.Start("MyServer", "4G")
```
