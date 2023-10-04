package zelligo

import (
	"fmt"
)

type ZellijPlugin interface {
	Load(configuration map[string]string) error
	Update(event Event) (bool, error)
	Render(rows uint32, cols uint32) error
}

var (
	STATE ZellijPlugin
)

func RegisterPlugin(p ZellijPlugin) {
	STATE = p
}

//export plugin_version
func pluginVersion() {
	fmt.Println("0.38.0")
}

//export load
func load() {
	pluginConfiguration := PluginConfiguration{}
	err := objectFromStdin(&pluginConfiguration)
	if err != nil {
		reportPanic(err)
		return
	}

	configuration := make(map[string]string)
	nameAndValues := pluginConfiguration.GetNameAndValue()
	if nameAndValues != nil {
		for _, g := range nameAndValues {
			if g != nil {
				configuration[g.Name] = g.Value
			}
		}
	}

	err = STATE.Load(configuration)
	if err != nil {
		reportPanic(err)
		return
	}
}

//export update
func update() bool {
	event := Event{}
	err := objectFromStdin(&event)
	if err != nil {
		reportPanic(err)
		return false
	}

	ret, err := STATE.Update(event)
	if err != nil {
		reportPanic(err)
		return false
	}
	return ret
}

//export render
func render(x, y uint32) {
	err := STATE.Render(x, y)
	if err != nil {
		reportPanic(err)
	}
}
