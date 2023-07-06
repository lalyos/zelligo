package gitlab.com/scabala/zelligo

import (
	"bufio"
	"os"
	"encoding/json"
	"runtime/debug"
	"fmt"
)

func objectToStdout(obj interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func objectFromStdin(obj interface{}) error {
	reader := bufio.NewReader(os.Stdin)

	data, err := reader.ReadBytes('\n')
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}

	return nil
}

//go:wasm-module zellij
//export host_subscribe
func hostSubscribe()

func Subscribe(eventTypes []EventType) error {
	err := objectToStdout(eventTypes)
	if err != nil {
		return nil
	}
	hostSubscribe()
	return nil
}

//go:wasm-module zellij
//export host_unsubscribe
func hostUnsubscribe()

func Unsubscribe(eventTypes []EventType) error {
	err := objectToStdout(eventTypes)
	if err != nil {
		return nil
	}
	hostUnsubscribe()
	return nil
}

//go:wasm-module zellij
//export set_selectable
func hostSetSelectable(selectable int32)

func SetSelectable(selectable bool) {
	var intSelectable int32
	if selectable {
		intSelectable = 1
	} else {
		intSelectable = 0
	}
	hostSetSelectable(intSelectable)
}

//go:wasm-module zellij
//export host_get_plugin_ids
func hostGetPluginIds()

func GetPluginIds() string {
	hostGetPluginIds()

	reader := bufio.NewReader(os.Stdin)
	plugin_ids, _ := reader.ReadString('\n')
	return plugin_ids
}

//go:wasm-module zellij
//export host_get_zellij_version
func hostGetZellijVersion()

func GetZellijVersion() (string, error) {
	hostGetZellijVersion()

	zellijVersion := ""
	err := objectFromStdin(&zellijVersion)
	if err != nil {
		return "", err
	}
	return zellijVersion, nil
}

//go:wasm-module zellij
//export host_open_file
func hostOpenFile()

func OpenFile(filepath string) error {
	err := objectToStdout(filepath)
	if err != nil {
		return err
	}
	hostOpenFile()
	return nil
}

//go:wasm-module zellij
//export host_open_file_floating
func hostOpenFileFloating()

func OpenFileFloating(filepath string) error {
	err := objectToStdout(filepath)
	if err != nil {
		return err
	}

	hostOpenFileFloating()
	return nil
}

//go:wasm-module zellij
//export host_open_file_with_line
func hostOpenFileWithLine()

func OpenFileWithLine(filepath string, line uint32) error {
	data := []interface{}{filepath, line}
	err := objectToStdout(data)
	if err != nil {
		return err
	}

	hostOpenFileWithLine()
	return nil
}

//go:wasm-module zellij
//export host_open_file_with_line_floating
func hostOpenFileWithLineFloating()

func OpenFileWithLineFloating(filepath string, line uint32) error {
	data := []interface{}{filepath, line}
	err := objectToStdout(data)
	if err != nil {
		return err
	}

	hostOpenFileWithLineFloating()
	return nil
}

//go:wasm-module zellij
//export host_open_terminal
func hostOpenTerminal()

func OpenTerminal(filepath string) error {
	err := objectToStdout(filepath)
	if err != nil {
		return err
	}

	hostOpenTerminal()
	return nil
}

//go:wasm-module zellij
//export host_open_terminal_floating
func hostOpenTerminalFloating()

func OpenTerminalFloating(filepath string) error {
	err := objectToStdout(filepath)
	if err != nil {
		return err
	}

	hostOpenTerminalFloating()
	return nil
}

//export host_open_command_pane
func hostOpenCommandPane()

func OpenCommandPane(command string, args []string) error {
	fullCommand := append([]string{command}, args...)
	err := objectToStdout(fullCommand)
	if err != nil {
		return err
	}
	hostOpenCommandPane()
	return nil
}

//export host_open_command_pane_floating
func hostOpenCommandPaneFloating()

func OpenCommandPaneFloating(command string, args []string) error {
	fullCommand := append([]string{command}, args...)
	err := objectToStdout(fullCommand)
	if err != nil {
		return err
	}
	hostOpenCommandPaneFloating()
	return nil
}

//go:wasm-module zellij
//export host_switch_to_tab
func hostSwitchToTab(tab_idx uint32)

func SwitchToTab(tab_idx uint32) error {
	hostSwitchToTab(tab_idx)
	return nil
}

//go:wasm-module zellij
//export host_set_timeout
func hostSetTimeout(secs float64)

func SetTimeout(secs float64) error {
	hostSetTimeout(secs)
	return nil
}

//go:wasm-module zellij
//export host_exec_cmd
func hostExecCmd()

func ExecCmd(cmd []string) error {
	err := objectToStdout(cmd)
	if err != nil {
		return err
	}
	hostExecCmd()
	return nil
}

//go:wasm-module zellij
//export host_hide_self
func hostHideSelf()

func HideSelf() {
	hostHideSelf()
}

//go:wasm-module zellij
//export host_show_self
func hostShowSelf(shouldFloatIfHidden uint32)

func ShowSelf(shouldFloatIfHidden bool) {
	var flag uint32
	if shouldFloatIfHidden {
		flag = 1
	} else {
		flag = 0
	}
	hostShowSelf(flag)
}

//go:wasm-module zellij
//export host_switch_to_input_mode
func hostSwitchToInputMode()

func SwitchToInputMode(inputMode InputMode) error {
	err := objectToStdout(inputMode)
	if err != nil {
		return err
	}
	hostSwitchToInputMode()
	return nil
}

//go:wasm-module zellij
//export host_new_tabs_with_layout
func hostNewTabsWithLayout()

func NewTabsWithLayout(layout string) error {
	err := objectToStdout(layout)
	if err != nil {
		return err
	}
	hostNewTabsWithLayout()
	return nil
}

//go:wasm-module zellij
//export host_new_tab
func hostNewTab()

func NewTab() {
	hostNewTab()
}

//go:wasm-module zellij
//export host_go_to_next_tab
func hostGoToNextTab()

func GoToNextTab() {
	hostGoToNextTab()
}

//go:wasm-module zellij
//export host_go_to_previous_tab
func hostGoToPreviousTab()

func GoToPreviousTab() {
	hostGoToPreviousTab()
}

//go:wasm-module zellij
//export host_report_panic
func hostReportPanic()

func reportPanic() {
	r := recover()
	if r != nil {
		stack := debug.Stack()
		stacktrace := fmt.Sprintf("%w\n\n%s", r, stack)

		fmt.Println("")
		fmt.Println("A panic occured in a plugin")
		fmt.Println(stacktrace)
		hostReportPanic()
	}
}


//TODO
//go:wasm-module zellij
//export host_resize_focused_pane
func hostResizeFocusedPane()

func ResizeFocusedPane(resize map[string]string) error {
	err := objectToStdout(resize)
	if err != nil {
		return err
	}
	hostResizeFocusedPane()
	return nil
}

//go:wasm-module zellij
//export host_resize_focused_pane_with_direction
func hostResizeFocusedPaneWithDirection()

func ResizeFocusedPaneWithDirection(resize map[string]string, direction map[string]string) error {
	object := []interface{}{resize, direction}
	err := objectToStdout(object)
	if err != nil {
		return err
	}
	hostResizeFocusedPaneWithDirection()
	return nil
}

//go:wasm-module zellij
//export host_focus_next_pan
func hostFocusNextPane()

func FocusNextPane() {
	hostFocusNextPane()
}

//go:wasm-module zellij
//export host_focus_previous_pan
func hostFocusPreviousPane()

func FocusPreviousPane() {
	hostFocusPreviousPane()
}

//TODO
//go:wasm-module zellij
//export host_move_focus
func hostMoveFocus()

func MoveFocus(direction map[string]string) error {
	err := objectToStdout(direction)
	if err != nil {
		return err
	}
	hostMoveFocus()
	return nil
}

//TODO
//go:wasm-module zellij
//export host_move_focus_or_tab
func hostMoveFocusOrTab()

func MoveFocusOrTab(direction map[string]string) error {
	err := objectToStdout(direction)
	if err != nil {
		return err
	}
	hostMoveFocusOrTab()
	return nil
}

//go:wasm-module zellij
//export host_detach
func hostDetach()

func Detach() {
	hostDetach()
}

//go:wasm-module zellij
//export host_edit_scrollback
func hostEditScrollback()

func EditScrollback() {
	hostEditScrollback()
}

//go:wasm-module zellij
//export host_write
func hostWrite()

func Write(bytes []byte) error {
	err := objectToStdout(bytes)
	if err != nil {
		return err
	}
	hostWrite()
	return nil
}

//go:wasm-module zellij
//export host_write_chars
func hostWriteChars()

func WriteChars(chars string) error {
	err := objectToStdout(chars)
	if err != nil {
		return err
	}
	hostWriteChars()
	return nil
}

//go:wasm-module zellij
//export host_toggle_tab
func hostToggleTab()

func ToggleTab() {
	hostToggleTab()
}

//go:wasm-module zellij
//export host_move_pane
func hostMovePane()

func MovePane() {
	hostMovePane()
}

//TODO
//go:wasm-module zellij
//export host_move_pane_with_direction
func hostMovePaneWithDirection()

func MovePaneWithDirection(direction map[string]string) error {
	err := objectToStdout(direction)
	if err != nil {
		return err
	}
	hostMovePaneWithDirection()
	return nil
}

//go:wasm-module zellij
//export host_clear_screen
func hostClearScreen()

func ClearScreen() {
	hostClearScreen()
}

//go:wasm-module zellij
//export host_scroll_up
func hostScrollUp()

func ScrollUp() {
	hostScrollUp()
}

//go:wasm-module zellij
//export host_scroll_down
func hostScrollDown()

func ScrollDown() {
	hostScrollDown()
}

//go:wasm-module zellij
//export host_scroll_to_top
func hostScrollToTop()

func ScrollToTop() {
	hostScrollToTop()
}

//go:wasm-module zellij
//export host_scroll_to_bottom
func hostScrollToBottom()

func ScrollToBottom() {
	hostScrollToBottom()
}

//go:wasm-module zellij
//export host_page_scroll_up
func hostPageScrollUp()

func PageScrollUp() {
	hostPageScrollUp()
}

//go:wasm-module zellij
//export host_page_scroll_down
func hostPageScrollDown()

func PageScrollDown() {
	hostPageScrollDown()
}

//go:wasm-module zellij
//export host_toggle_focus_fullscreen
func hostToggleFocusFullscreen()

func ToggleFocusFullscreen() {
	hostToggleFocusFullscreen()
}

//go:wasm-module zellij
//export host_toggle_pane_frames
func hostTogglePaneFrames()

func TogglePaneFrames() {
	hostTogglePaneFrames()
}

//go:wasm-module zellij
//export host_toggle_pane_embed_or_eject
func hostTogglePaneEmbedOrEject()

func TogglePaneEmbedOrEject() {
	hostTogglePaneEmbedOrEject()
}

//go:wasm-module zellij
//export host_undo_rename_pane
func hostUndoRenamePane()

func UndoRenamePane() {
	hostUndoRenamePane()
}

//go:wasm-module zellij
//export host_close_focus
func hostCloseFocus()

func CloseFocus() {
	hostCloseFocus()
}

//go:wasm-module zellij
//export host_toggle_active_tab_sync
func hostToggleActiveTabSync()

func ToggleActiveTabSync() {
	hostToggleActiveTabSync()
}

//go:wasm-module zellij
//export host_close_focused_tab
func hostCloseFocusedTab()

func CloseFocusedTab() {
	hostCloseFocusedTab()
}

//go:wasm-module zellij
//export host_undo_rename_tab
func hostUndoRenameTab()

func UndoRenameTab() {
	hostUndoRenameTab()
}

//go:wasm-module zellij
//export host_quit_zellij
func hostQuitZellij()

func QuitZellij() {
	hostQuitZellij()
}

//go:wasm-module zellij
//export host_previous_swap_layout
func hostPreviousSwapLayout()

func PreviousSwapLayout() {
	hostPreviousSwapLayout()
}

//go:wasm-module zellij
//export host_next_swap_layout
func hostNextSwapLayout()

func NextSwapLayout() {
	hostNextSwapLayout()
}

//go:wasm-module zellij
//export host_go_to_tab_name
func hostGoToTabName()

func GoToTabName(tabName string) error {
	err := objectToStdout(tabName)
	if err != nil {
		return err
	}
	hostGoToTabName()
	return nil
}

//go:wasm-module zellij
//export host_focus_or_create_tab
func hostFocusOrCreateTab()

func FocusOrCreateTab(tabName string) error {
	err := objectToStdout(tabName)
	if err != nil {
		return err
	}
	hostFocusOrCreateTab()
	return nil
}

//go:wasm-module zellij
//export host_go_to_tab
func hostGoToTab(tabIndex uint32)

func GoToTab(tabIndex uint32) {
	hostGoToTab(tabIndex)
}

//go:wasm-module zellij
//export host_start_or_reload_plugin
func hostStartOrReloadPlugin()

func StartOrReloadPlugin(url string) error {
	err := objectToStdout(url)
	if err != nil {
		return err
	}
	hostStartOrReloadPlugin()
	return nil
}

//go:wasm-module zellij
//export host_close_terminal_pane
func hostCloseTerminalPane(terminalPaneId uint32)

func CloseTerminalPane(terminalPaneId uint32) {
	hostCloseTerminalPane(terminalPaneId)
}

//go:wasm-module zellij
//export host_close_terminal_pane
func hostClosePluginPane(pluginPaneId uint32)

func ClosePluginPane(pluginPaneId uint32) {
	hostClosePluginPane(pluginPaneId)
}

//go:wasm-module zellij
//export host_focus_terminal_pane
func hostFocusTerminalPane(terminalPaneId int32, shouldFloatIfHidden int32)

func FocusTerminalPane(terminalPaneId int32, shouldFloatIfHidden bool) {
	var shouldFloat int32
	if shouldFloatIfHidden {
		shouldFloat = 1
	} else {
		shouldFloat = 0
	}
	hostFocusTerminalPane(terminalPaneId, shouldFloat)
}

//go:wasm-module zellij
//export host_focus_plugin_pane
func hostFocusPluginPane(pluginPaneId int32, shouldFloatIfHidden int32)

func CloseFocusPane(pluginPaneId int32, shouldFloatIfHidden bool) {
	var shouldFloat int32
	if shouldFloatIfHidden {
		shouldFloat = 1
	} else {
		shouldFloat = 0
	}
	hostFocusPluginPane(pluginPaneId, shouldFloat)
}

//go:wasm-module zellij
//export host_rename_terminal_pane
func hostRenameTerminalPane()

func RenameTerminalPane(terminalPaneId uint32, newName string) error {
	object := []interface{}{terminalPaneId, newName}
	err := objectToStdout(object)
	if err != nil {
		return err
	}
	hostRenameTerminalPane()
	return nil
}

//go:wasm-module zellij
//export host_rename_plugin_pane
func hostRenamePluginPane()

func RenamePluginPane(pluginPaneId uint32, newName string) error {
	object := []interface{}{pluginPaneId, newName}
	err := objectToStdout(object)
	if err != nil {
		return err
	}
	hostRenamePluginPane()
	return nil
}

//go:wasm-module zellij
//export host_rename_tab
func hostRenameTab()

func RenameTab(tabPosition uint32, newName string) error {
	object := []interface{}{tabPosition, newName}
	err := objectToStdout(object)
	if err != nil {
		return err
	}
	hostRenamePluginPane()
	return nil
}

//go:wasm-module zellij
//export host_post_message_to
func hostPostMessageTo()

func PostMessageTo(workerName string, message string, payload string) error {
	object := []interface{}{workerName, message, payload}
	err := objectToStdout(object)
	if err != nil {
		return err
	}
	hostPostMessageTo()
	return nil
}

//go:wasm-module zellij
//export host_post_message_to_plugin
func hostPostMessageToPlugin()

func PostMessageToPlugin(message string, payload string) error {
	object := []interface{}{message, payload}
	err := objectToStdout(object)
	if err != nil {
		return err
	}
	hostPostMessageToPlugin()
	return nil
}

