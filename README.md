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

Main target for compiling Go to WASI is [tinygo](https://tinygo.org/) and it is the compiler used with this project.

In order to build plugin, use following command:

```bash
tinygo build -o plugin.wasm -target=wasi
```

It will produce plugin `plugin.wasm` in current directory, using Go sources from current directory.

One could use also currently (1.21) experimental port of Go to WASI preview 1 by running:

```bash
env GOOS=wasip1 GOARCH=wasm go build -o plugin.wasm main.go
```

However, it is young and experimental so expect bugs.

## Version support

zelligo supports Zellij version 0.38.0 and later.

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


func (p *MyPlugin) Load(configuration map[string]string) {
	fmt.Fprintf(os.Stderr, "configuration for plugin is: %v", configuration)
	zelligo.Subscribe([]zelligo.EventType{
		zelligo.EventTypeKey,
	})
}

func (p *MyPlugin) Update(event map[string]interface{}) bool {
	fmt.Fprintf(os.Stderr, "log message: %v", event)
	return true
}

func (p *MyPlugin) Render(cols uint32, rows uint32) {
	fmt.Println(p.message)
}

func init() {
	zelligo.RegisterPlugin(&MyPlugin{message: "Hello from Go!"})
}

func main() {}  // this is very important
```

Please keep in mind that you **have to** define main function. It should be empty - 
becasue it is not used - but must be defined to properly compile to WASI.

## Development

In order to develop `zelligo`, you need following:

- Go version 1.21 or up/tinygo
- [protoc](https://github.com/protocolbuffers/protobuf#protobuf-compiler-installation)
- [protoc-gen-go](https://github.com/golang/protobuf) protoc plugin
- [protoc-gen-go-vtproto](https://github.com/planetscale/vtprotobuf) vtprotobuf protoc plugin

## Improvements (MRs welcomed!)

 - Workers support

