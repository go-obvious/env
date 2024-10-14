# Obvious `env` helper methods

Simple environment variable helper library

## How to Use


### Installation

```sh
go get github.com/go-obvious/env
```

### Example Usage

```go
package main

import (
    "fmt"
    "github.com/go-obvious/env"
)

func main() {
    // Load environment variables from a .env file
    value := env.GetOr("MY_VARIABLE", "nothing")
    fmt.Println("Got:", value)
}
```

### License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
