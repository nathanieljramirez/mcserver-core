# MCServerKit

A Go library for creating and managing Minecraft servers.

## Installation

```bash
go get mcserverkit.github.io
```

## Usage

Install a server version

```go
mcserverkit.Install(version string)
```

Create a server

```go
mcserverkit.Create(name string, eula bool)
```

Start your server

```go
mcserverkit.Start(name string, memory ...string)
```

## Example

```go
package main

import (
	"fmt"

	"mcserverkit.github.io"
)

func main() {
	err := mcserverkit.Install("1.21.1")
	if err != nil {
		fmt.Println("Error installing 1.21.1:", err)
		return
	}

	err = mcserverkit.Create("MyServer", true)
	if err != nil {
		fmt.Println("Error creating MyServer:", err)
		return
	}

	err = mcserverkit.Start("MyServer", "4G")
	if err != nil {
		fmt.Println("Error starting MyServer:", err)
		return
	}
}
```
