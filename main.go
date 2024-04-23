package zelligo

import (
	"bytes"
	"fmt"

	kdl "github.com/sblinch/kdl-go"
)

type ZellijPlugin interface {
	Load(configuration []byte) error
	Update(event Event) (bool, error)
	Pipe(message PipeMessage) (bool, error)
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

	nameAndValues := pluginConfiguration.GetNameAndValue()
	if nameAndValues == nil {
		err = STATE.Load([]byte{})
		if err != nil {
			reportPanic(err)
		}
		return
	}

	configMap := make(map[string]string)
	for _, g := range nameAndValues {
		if g != nil {
			configMap[g.Name] = g.Value
		}
	}

	configBuffer := bytes.Buffer{}

	enc := kdl.NewEncoder(&configBuffer)
	enc.Encode(configMap)
	if err != nil {
		reportPanic(fmt.Errorf("cannot encode configuration into KDL: %v", err))
		return
	}

	err = STATE.Load(configBuffer.Bytes())
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

//export pipe
func pipe() bool {
	message := PipeMessage{}
	err := objectFromStdin(&message)
	if err != nil {
		reportPanic(err)
		return false
	}

	ret, err := STATE.Pipe(message)
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
