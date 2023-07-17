package zelligo

import (
	"fmt"
	"bufio"
	"os"
)

type ZellijPlugin interface {
	Load()
	Update(event map[string]string) bool
	Render(rows uint32, cols uint32)
}

var (
	STATE ZellijPlugin
)

func RegisterPlugin(p ZellijPlugin) {
	STATE = p
}

//export plugin_version
func pluginVersion() {
	fmt.Println("0.37.2")
}

//export load
func load() {
	defer reportPanic()
	STATE.Load()
}

//export update
func update() bool {
	defer reportPanic()

	event := make(map[string]string)
	err := objectfromstdin(&event)
	if err != nil {
		panic(err)
	}

	ret := STATE.Update(event)
	return ret
}

//export render
func render(x, y uint32) {
	defer reportPanic()
	STATE.Render(x, y)
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}

