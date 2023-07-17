# zelligo - Zellij Go SDK

[zellij](https://zellij.dev/) is modern terminal multiplexer written in Rust with WebAssembly plugin system. It has native support for creating plugins in Rust. This project strives to make creating Zellij plugins in Go as easy as in Rust.

WARNING! - this project is on very early stage of development! Expect bugs and breaking changes!

## What is zelligo

[zelligo]() is a Go library that exposes nice API for writing Zellij plugins. It wraps WebAssembly internals and data structures encoding and decoding behind `ZellijPlugin` interface.
In order to create a plugin, just implement this interface and register plugin with `RegisterPlugin` function and done!

For more detailed documentation, refer to [official docs](https://zellij.dev/documentation/plugins.html)

## Error handling

As is advised by Go, [you should handle errors](https://go.dev/doc/tutorial/handle-errors). When encoutered an error that you cannot handle, just panic :) - zelligo catches panics and reports them to Zellij.

## Building

Current Go version (1.20) does not support building WASI binaries and only Go compiler that supports building WASI binaries is [tinygo](https://tinygo.org/).

In order to build plugin, use following command:

```bash
tinygo build -o plugin.wasm -target=wasi
```

It will produce plugin `plugin.wasm` in current directory, using Go sources from current directory.

## Version support

zelligo supports Zellij version 0.37.0 and later.

## Example usage

```go
package main

import (
	"fmt"
	"os"

	"gitlab.com/scabala/zelligo"
)

type MyPlugin struct {
	message string
}


func (p *MyPlugin) Load() {
	zelligo.Subscribe([]zelligo.EventType{
		zelligo.EventTypeKey,
	})
}

func (p *MyPlugin) Update(event map[string]string) bool {
	fmt.Fprintf(os.Stderr, "log message: %v", event)
	return true
}

func (p *MyPlugin) Render(cols uint32, rows uint32) {
	fmt.Println(p.message)
}

func init() {
	zelligo.RegisterPlugin(&MyPlugin{message: "Hello from Go!"})
}
```

## Improvements (MRs welcomed!)

 - Workers support
 - Custom structures instead of `map[string]string`

