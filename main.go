package zelligo

import (
	"fmt"
)

type ZellijPlugin interface {
	Load()
	Update(event map[string]interface{}) bool
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

	event := make(map[string]interface{})
	err := objectFromStdin(&event)
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
