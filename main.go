package zelligo

import (
	"fmt"
)

type ZellijPlugin interface {
	Load(configuration map[string]string)
	Update(event Event) bool
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

	pluginConfiguration := PluginConfiguration{}
	err := objectFromStdin(&pluginConfiguration)
	if err != nil {
		panic(err)
	}

	configuration := make(map[string]string)
	for _, g := range pluginConfiguration.NameAndValue {
		if g != nil {
			configuration[g.Name] = g.Value
		}
	}

	STATE.Load(configuration)
}

//export update
func update() bool {
	defer reportPanic()

	event := Event{}
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
