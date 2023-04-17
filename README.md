# wasm-tinygo-pdk

## Write a wasm plugin callable by Node.js

```golang
package main

import (
	"fmt"
	"os"
	"path/filepath"

	plugin "github.com/bots-garden/wasm-tinygo-pdk"
)

func main() {
	plugin.SetHandle(Handle)
}

func Handle(param []byte) ([]byte, error) {

	return []byte("Hello " + string(param)), nil
}
```
