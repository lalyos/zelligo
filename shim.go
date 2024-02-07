package zelligo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	kdl "github.com/sblinch/kdl-go"
)

type Serializable interface {
	MarshalVT() ([]byte, error)
	UnmarshalVT([]byte) error
}

func objectToStdout(obj Serializable) error {
	data, err := obj.MarshalVT()
	if err != nil {
		return err
	}

	/*
		https://pkg.go.dev/encoding/json#Marshal

		> Array and slice values encode as JSON arrays, except that []byte encodes as a base64-encoded string, and a nil slice encodes as the null JSON value.

		Therefore, we need to manually encode []byte to a JSON array of integers
	*/
	/*encodedData, err := json.Marshal(data)
	if err != nil {
		return err
	}*/
	encodedData := strings.Join(strings.Fields(fmt.Sprintf("%d", data)), ",")

	fmt.Println(encodedData)

	return nil
}

func objectFromStdin(obj Serializable) error {
	reader := bufio.NewReader(os.Stdin)

	encodedData, err := reader.ReadBytes('\n')
	if err != nil {
		return err
	}

	var data []byte
	err = json.Unmarshal(encodedData, &data)
	if err != nil {
		return err
	}

	err = obj.UnmarshalVT(data)
	if err != nil {
		return err
	}

	return nil
}

//go:wasmimport zellij host_run_plugin_command
func hostRunPluginCommand()

func Subscribe(eventTypes []EventType) error {
	pc := PluginCommand{
		Name: CommandName_Subscribe,
		Payload: &PluginCommand_SubscribePayload{
			SubscribePayload: &SubscribePayload{
				Subscriptions: &EventNameList{
					EventTypes: eventTypes,
				},
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func Unsubscribe(eventTypes []EventType) error {
	pc := PluginCommand{
		Name: CommandName_Unsubscribe,
		Payload: &PluginCommand_UnsubscribePayload{
			UnsubscribePayload: &UnsubscribePayload{
				Subscriptions: &EventNameList{
					EventTypes: eventTypes,
				},
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func SetSelectable(selectable bool) error {
	pc := PluginCommand{
		Name: CommandName_Unsubscribe,
		Payload: &PluginCommand_SetSelectablePayload{
			SetSelectablePayload: selectable,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func RequestPermission(permissions []PermissionType) error {
	pc := PluginCommand{
		Name: CommandName_RequestPluginPermissions,
		Payload: &PluginCommand_RequestPluginPermissionPayload{
			RequestPluginPermissionPayload: &RequestPluginPermissionPayload{
				Permissions: permissions,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func GetPluginIds() (string, error) {
	pc := PluginCommand{
		Name:    CommandName_GetPluginIds,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return "", err
	}

	hostRunPluginCommand()

	pluginIDs := PluginIds{}
	err = objectFromStdin(&pluginIDs)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(pluginIDs.GetPluginId()), nil
}

func GetZellijVersion() (string, error) {
	pc := PluginCommand{
		Name:    CommandName_GetZellijVersion,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return "", err
	}

	hostRunPluginCommand()

	zellijVersion := ZellijVersion{}
	err = objectFromStdin(&zellijVersion)
	if err != nil {
		return "", err
	}
	return zellijVersion.GetVersion(), nil
}

func OpenFile(file *File) error {
	pc := PluginCommand{
		Name: CommandName_OpenFile,
		Payload: &PluginCommand_OpenFilePayload{
			OpenFilePayload: &OpenFilePayload{
				FileToOpen: file,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func OpenFileFloating(file *File) error {
	pc := PluginCommand{
		Name: CommandName_OpenFileFloating,
		Payload: &PluginCommand_OpenFileFloatingPayload{
			OpenFileFloatingPayload: &OpenFilePayload{
				FileToOpen: file,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func OpenFileInPlace(file *File) error {
	pc := PluginCommand{
		Name: CommandName_OpenFileInPlace,
		Payload: &PluginCommand_OpenFileInPlacePayload{
			OpenFileInPlacePayload: &OpenFilePayload{
				FileToOpen: file,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func OpenTerminal(filepath string) error {
	pc := PluginCommand{
		Name: CommandName_OpenTerminal,
		Payload: &PluginCommand_OpenTerminalPayload{
			OpenTerminalPayload: &OpenFilePayload{
				FileToOpen: &File{
					Path: filepath,
				},
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func OpenTerminalFloating(filepath string) error {
	pc := PluginCommand{
		Name: CommandName_OpenTerminalFloating,
		Payload: &PluginCommand_OpenTerminalFloatingPayload{
			OpenTerminalFloatingPayload: &OpenFilePayload{
				FileToOpen: &File{
					Path: filepath,
				},
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func OpenTerminalInPlace(filepath string) error {
	pc := PluginCommand{
		Name: CommandName_OpenTerminalInPlace,
		Payload: &PluginCommand_OpenTerminalInPlacePayload{
			OpenTerminalInPlacePayload: &OpenFilePayload{
				FileToOpen: &File{
					Path: filepath,
				},
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func OpenCommandPane(command *Command) error {
	pc := PluginCommand{
		Name: CommandName_OpenCommandPane,
		Payload: &PluginCommand_OpenCommandPanePayload{
			OpenCommandPanePayload: &OpenCommandPanePayload{
				CommandToRun: command,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func OpenCommandPaneFloating(command *Command) error {
	pc := PluginCommand{
		Name: CommandName_OpenCommandPaneFloating,
		Payload: &PluginCommand_OpenCommandPaneFloatingPayload{
			OpenCommandPaneFloatingPayload: &OpenCommandPanePayload{
				CommandToRun: command,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func OpenCommandPaneInPlace(command *Command) error {
	pc := PluginCommand{
		Name: CommandName_OpenCommandInPlace,
		Payload: &PluginCommand_OpenCommandPaneInPlacePayload{
			OpenCommandPaneInPlacePayload: &OpenCommandPanePayload{
				CommandToRun: command,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func SwitchToTab(tabIdx uint32) error {
	pc := PluginCommand{
		Name: CommandName_SwitchTabTo,
		Payload: &PluginCommand_SwitchTabToPayload{
			SwitchTabToPayload: &SwitchTabToPayload{
				TabIndex: tabIdx,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func SetTimeout(secs float64) error {
	pc := PluginCommand{
		Name: CommandName_SetTimeout,
		Payload: &PluginCommand_SetTimeoutPayload{
			SetTimeoutPayload: &SetTimeoutPayload{
				Seconds: secs,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ExecCmd(cmd []string) error {
	pc := PluginCommand{
		Name: CommandName_ExecCmd,
		Payload: &PluginCommand_ExecCmdPayload{
			ExecCmdPayload: &ExecCmdPayload{
				CommandLine: cmd,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

type RunCommandOptions struct {
	WorkDir string
	EnvVars map[string]string
	Context map[string]string
}

func RunCommand(cmd []string, opt *RunCommandOptions) error {
	//default option values
	env := make([]*EnvVariable, 0)
	ctx := make([]*ContextItem, 0)
	cwd := "."

	if opt != nil && opt.Context != nil {
		for k, v := range opt.Context {
			ctx = append(ctx, &ContextItem{Name: k, Value: v})
		}
	}

	if opt != nil && opt.EnvVars != nil {
		for k, v := range opt.EnvVars {
			env = append(env, &EnvVariable{Name: k, Value: v})
		}
	}

	if opt != nil && opt.WorkDir != "" {
		cwd = opt.WorkDir
	}

	pc := PluginCommand{
		Name: CommandName_RunCommand,
		Payload: &PluginCommand_RunCommandPayload{
			RunCommandPayload: &RunCommandPayload{
				CommandLine:  cmd,
				EnvVariables: env,
				Cwd:          cwd,
				Context:      ctx,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

type WebRequestOptions struct {
	Verb    string
	Headers map[string]string
	Body    []byte
	Context map[string]string
}

func WebRequest(url string, opt *WebRequestOptions) error {
	//default option values
	verb := HttpVerb_Get
	headers := make([]*Header, 0)
	body := []byte{}
	ctx := make([]*ContextItem, 0)

	if opt != nil && opt.Context != nil {
		for k, v := range opt.Context {
			ctx = append(ctx, &ContextItem{Name: k, Value: v})
		}
	}

	if opt != nil && opt.Headers != nil {
		for k, v := range opt.Headers {
			headers = append(headers, &Header{Name: k, Value: v})
		}
	}

	if opt != nil && opt.Verb != "" {
		v, found := HttpVerb_value[opt.Verb]
		if !found {
			return nil
		}
		verb = HttpVerb(v)
	}

	pc := PluginCommand{
		Name: CommandName_WebRequest,
		Payload: &PluginCommand_WebRequestPayload{
			WebRequestPayload: &WebRequestPayload{
				Url:     url,
				Verb:    verb,
				Headers: headers,
				Body:    body,
				Context: ctx,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func HideSelf() error {
	pc := PluginCommand{
		Name:    CommandName_HideSelf,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ShowSelf(shouldFloatIfHidden bool) error {
	pc := PluginCommand{
		Name: CommandName_HideSelf,
		Payload: &PluginCommand_ShowSelfPayload{
			ShowSelfPayload: shouldFloatIfHidden,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func SwitchToInputMode(inputMode InputMode) error {
	pc := PluginCommand{
		Name: CommandName_SwitchToMode,
		Payload: &PluginCommand_SwitchToModePayload{
			SwitchToModePayload: &SwitchToModePayload{
				InputMode: inputMode,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func NewTabsWithLayout(layout string) error {
	pc := PluginCommand{
		Name: CommandName_NewTabsWithLayout,
		Payload: &PluginCommand_NewTabsWithLayoutPayload{
			NewTabsWithLayoutPayload: layout,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func NewTab() error {
	pc := PluginCommand{
		Name:    CommandName_NewTab,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func GoToNextTab() error {
	pc := PluginCommand{
		Name:    CommandName_GoToNextTab,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func GoToPreviousTab() error {
	pc := PluginCommand{
		Name:    CommandName_GoToPreviousTab,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func reportPanic(errToReport error) {
	stack := string(debug.Stack())
	stacktrace := fmt.Sprintf("%s\n\n%s", errToReport, stack)
	panicPayload := strings.ReplaceAll(stacktrace, "\n", "\r\n")

	pc := PluginCommand{
		Name: CommandName_ReportCrash,
		Payload: &PluginCommand_ReportCrashPayload{
			ReportCrashPayload: panicPayload,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could report panic. Dumping panic to stderr. Panic: %s\n", panicPayload)
	}

	hostRunPluginCommand()
}

func ResizeFocusedPane(resize ResizeAction) error {
	pc := PluginCommand{
		Name: CommandName_Resize,
		Payload: &PluginCommand_ResizePayload{
			ResizePayload: &ResizePayload{
				Resize: &Resize{
					ResizeAction: resize,
					Direction:    nil,
				},
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ResizeFocusedPaneWithDirection(resize ResizeAction, direction ResizeDirection) error {
	pc := PluginCommand{
		Name: CommandName_Resize,
		Payload: &PluginCommand_ResizePayload{
			ResizePayload: &ResizePayload{
				Resize: &Resize{
					ResizeAction: resize,
					Direction:    &direction,
				},
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func FocusNextPane() error {
	pc := PluginCommand{
		Name:    CommandName_FocusNextPane,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func FocusPreviousPane() error {
	pc := PluginCommand{
		Name:    CommandName_FocusPreviousPane,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func MoveFocus(direction ResizeDirection) error {
	pc := PluginCommand{
		Name: CommandName_MoveFocus,
		Payload: &PluginCommand_MoveFocusPayload{
			MoveFocusPayload: &MovePayload{
				Direction: &MoveDirection{
					Direction: direction,
				},
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func MoveFocusOrTab(direction ResizeDirection) error {
	pc := PluginCommand{
		Name: CommandName_MoveFocusOrTab,
		Payload: &PluginCommand_MoveFocusOrTabPayload{
			MoveFocusOrTabPayload: &MovePayload{
				Direction: &MoveDirection{
					Direction: direction,
				},
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func Detach() error {
	pc := PluginCommand{
		Name:    CommandName_Detach,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func EditScrollback() error {
	pc := PluginCommand{
		Name:    CommandName_EditScrollback,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func Write(bytes []byte) error {
	pc := PluginCommand{
		Name: CommandName_Write,
		Payload: &PluginCommand_WritePayload{
			WritePayload: bytes,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func WriteChars(chars string) error {
	pc := PluginCommand{
		Name: CommandName_WriteChars,
		Payload: &PluginCommand_WriteCharsPayload{
			WriteCharsPayload: chars,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ToggleTab() error {
	pc := PluginCommand{
		Name:    CommandName_ToggleTab,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func MovePane() error {
	pc := PluginCommand{
		Name:    CommandName_MovePane,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func MovePaneWithDirection(direction ResizeDirection) error {
	pc := PluginCommand{
		Name: CommandName_MovePaneWithDirection,
		Payload: &PluginCommand_MovePaneWithDirectionPayload{
			MovePaneWithDirectionPayload: &MovePayload{
				Direction: &MoveDirection{
					Direction: direction,
				},
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ClearScreen() error {
	pc := PluginCommand{
		Name:    CommandName_ClearScreen,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ScrollUp() error {
	pc := PluginCommand{
		Name:    CommandName_ScrollUp,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ScrollDown() error {
	pc := PluginCommand{
		Name:    CommandName_ScrollDown,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ScrollToTop() error {
	pc := PluginCommand{
		Name:    CommandName_ScrollToTop,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ScrollToBottom() error {
	pc := PluginCommand{
		Name:    CommandName_ScrollToBottom,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func PageScrollUp() error {
	pc := PluginCommand{
		Name:    CommandName_PageScrollUp,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func PageScrollDown() error {
	pc := PluginCommand{
		Name:    CommandName_PageScrollDown,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ToggleFocusFullscreen() error {
	pc := PluginCommand{
		Name:    CommandName_ToggleFocusFullscreen,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func TogglePaneFrames() error {
	pc := PluginCommand{
		Name:    CommandName_TogglePaneFrames,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func TogglePaneEmbedOrEject() error {
	pc := PluginCommand{
		Name:    CommandName_TogglePaneEmbedOrEject,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func UndoRenamePane() error {
	pc := PluginCommand{
		Name:    CommandName_UndoRenamePane,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func CloseFocus() error {
	pc := PluginCommand{
		Name:    CommandName_CloseFocus,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ToggleActiveTabSync() error {
	pc := PluginCommand{
		Name:    CommandName_ToggleActiveTabSync,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func CloseFocusedTab() error {
	pc := PluginCommand{
		Name:    CommandName_CloseFocusedTab,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func UndoRenameTab() error {
	pc := PluginCommand{
		Name:    CommandName_UndoRenameTab,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func QuitZellij() error {
	pc := PluginCommand{
		Name:    CommandName_QuitZellij,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func PreviousSwapLayout() error {
	pc := PluginCommand{
		Name:    CommandName_PreviousSwapLayout,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func NextSwapLayout() error {
	pc := PluginCommand{
		Name:    CommandName_NextSwapLayout,
		Payload: nil,
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func GoToTabName(tabName string) error {
	pc := PluginCommand{
		Name: CommandName_GoToTabName,
		Payload: &PluginCommand_GoToTabNamePayload{
			GoToTabNamePayload: tabName,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func FocusOrCreateTab(tabName string) error {
	pc := PluginCommand{
		Name: CommandName_FocusOrCreateTab,
		Payload: &PluginCommand_FocusOrCreateTabPayload{
			FocusOrCreateTabPayload: tabName,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func GoToTab(tabIndex uint32) error {
	pc := PluginCommand{
		Name: CommandName_GoToTab,
		Payload: &PluginCommand_GoToTabPayload{
			GoToTabPayload: tabIndex,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func StartOrReloadPlugin(url string) error {
	pc := PluginCommand{
		Name: CommandName_StartOrReloadPlugin,
		Payload: &PluginCommand_StartOrReloadPluginPayload{
			StartOrReloadPluginPayload: url,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func CloseTerminalPane(terminalPaneId uint32) error {
	pc := PluginCommand{
		Name: CommandName_CloseTerminalPane,
		Payload: &PluginCommand_CloseTerminalPanePayload{
			CloseTerminalPanePayload: terminalPaneId,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func ClosePluginPane(pluginPaneId uint32) error {
	pc := PluginCommand{
		Name: CommandName_ClosePluginPane,
		Payload: &PluginCommand_ClosePluginPanePayload{
			ClosePluginPanePayload: pluginPaneId,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func FocusTerminalPane(terminalPaneId uint32, shouldFloat bool) error {
	pc := PluginCommand{
		Name: CommandName_FocusTerminalPane,
		Payload: &PluginCommand_FocusTerminalPanePayload{
			FocusTerminalPanePayload: &PaneIdAndShouldFloat{
				PaneId:      terminalPaneId,
				ShouldFloat: shouldFloat,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func FocusPluginPane(pluginPaneId uint32, shouldFloat bool) error {
	pc := PluginCommand{
		Name: CommandName_FocusPluginPane,
		Payload: &PluginCommand_FocusPluginPanePayload{
			FocusPluginPanePayload: &PaneIdAndShouldFloat{
				PaneId:      pluginPaneId,
				ShouldFloat: shouldFloat,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func RenameTerminalPane(terminalPaneId uint32, newName string) error {
	pc := PluginCommand{
		Name: CommandName_RenameTerminalPane,
		Payload: &PluginCommand_RenameTerminalPanePayload{
			RenameTerminalPanePayload: &IdAndNewName{
				Id:      terminalPaneId,
				NewName: newName,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func RenamePluginPane(pluginPaneId uint32, newName string) error {
	pc := PluginCommand{
		Name: CommandName_RenamePluginPane,
		Payload: &PluginCommand_RenamePluginPanePayload{
			RenamePluginPanePayload: &IdAndNewName{
				Id:      pluginPaneId,
				NewName: newName,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func RenameTab(tabPosition uint32, newName string) error {
	pc := PluginCommand{
		Name: CommandName_RenameTab,
		Payload: &PluginCommand_RenameTabPayload{
			RenameTabPayload: &IdAndNewName{
				Id:      tabPosition,
				NewName: newName,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func SwitchSession(name *string) error {
	pc := PluginCommand{
		Name: CommandName_SwitchSession,
		Payload: &PluginCommand_SwitchSessionPayload{
			SwitchSessionPayload: &SwitchSessionPayload{
				Name: name,
			},
		},
	}

	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func SwitchSessionWithLayout(name *string, layoutInfo LayoutInfo) error {
	pc := PluginCommand{
		Name: CommandName_SwitchSession,
		Payload: &PluginCommand_SwitchSessionPayload{
			SwitchSessionPayload: &SwitchSessionPayload{
				Name:   name,
				Layout: &layoutInfo,
			},
		},
	}

	err := objectToStdout(&pc)
	if err != nil {
		return err
	}

	hostRunPluginCommand()
	return nil
}

func SwitchSessionWithFocus(name string, tabPosition *uint32, paneId *uint32) error {
	pc := PluginCommand{
		Name: CommandName_SwitchSession,
		Payload: &PluginCommand_SwitchSessionPayload{
			SwitchSessionPayload: &SwitchSessionPayload{
				Name:        &name,
				TabPosition: tabPosition,
				PaneId:      paneId,
			},
		},
	}

	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func DeleteDeadSession(name string) error {
	pc := PluginCommand{
		Name: CommandName_DeleteDeadSession,
		Payload: &PluginCommand_DeleteDeadSessionPayload{
			DeleteDeadSessionPayload: name,
		},
	}

	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func DeleteAllDeadSessions() error {
	pc := PluginCommand{
		Name:    CommandName_DeleteDeadSession,
		Payload: nil,
	}

	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func RenameSession(name string) error {
	pc := PluginCommand{
		Name: CommandName_RenameSession,
		Payload: &PluginCommand_RenameSessionPayload{
			RenameSessionPayload: name,
		},
	}

	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func UnblockCliPipeInput(pipe string) error {
	pc := PluginCommand{
		Name: CommandName_UnblockCliPipeInput,
		Payload: &PluginCommand_UnblockCliPipeInputPayload{
			UnblockCliPipeInputPayload: pipe,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func BlockCliPipeInput(pipe string) error {
	pc := PluginCommand{
		Name: CommandName_BlockCliPipeInput,
		Payload: &PluginCommand_BlockCliPipeInputPayload{
			BlockCliPipeInputPayload: pipe,
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func CliPipeOutput(pipe string, output string) error {
	pc := PluginCommand{
		Name: CommandName_CliPipeOutput,
		Payload: &PluginCommand_CliPipeOutputPayload{
			CliPipeOutputPayload: &CliPipeOutputPayload{
				PipeName: pipe,
				Output:   output,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

type PipeMessage struct {
	Name    string
	Payload *string
	Args    map[string]string

	PluginUrl     *string
	PluginConfig  map[string]interface{}
	NewPluginArgs *NewPluginArgs
}

func PipeMessageToPlugin(message *PipeMessage) error {
	messageArgs := make([]*ContextItem, 0)
	pluginConfig := make([]*ContextItem, 0)

	if message.Args != nil {
		for k, v := range message.Args {
			messageArgs = append(messageArgs, &ContextItem{Name: k, Value: v})
		}
	}

	if message.PluginUrl != nil && message.PluginConfig != nil {
		for k, v := range message.PluginConfig {
			kdlValue, err := kdl.Marshal(v)
			if err != nil {
				return err
			}
			pluginConfig = append(pluginConfig, &ContextItem{Name: k, Value: string(kdlValue)})
		}
	}

	pc := PluginCommand{
		Name: CommandName_MessageToPlugin,
		Payload: &PluginCommand_MessageToPluginPayload{
			MessageToPluginPayload: &MessageToPluginPayload{
				PluginUrl:      message.PluginUrl,
				PluginConfig:   pluginConfig,
				MessageName:    message.Name,
				MessagePayload: message.Payload,
				MessageArgs:    messageArgs,
				NewPluginArgs:  message.NewPluginArgs,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func PostMessageTo(message *Message) error {
	pc := PluginCommand{
		Name: CommandName_PostMessageTo,
		Payload: &PluginCommand_PostMessageToPayload{
			PostMessageToPayload: &PluginMessagePayload{
				Message: message,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}

func PostMessageToPlugin(message *Message) error {
	pc := PluginCommand{
		Name: CommandName_PostMessageToPlugin,
		Payload: &PluginCommand_PostMessageToPluginPayload{
			PostMessageToPluginPayload: &PluginMessagePayload{
				Message: message,
			},
		},
	}
	err := objectToStdout(&pc)
	if err != nil {
		return err
	}
	hostRunPluginCommand()
	return nil
}
