package gitlab.com/scabala/zelligo

type InputMode string

const (
	InputModeNormal InputMode = "normal"
	InputModeLocked InputMode = "locked"
	InputModeResize InputMode = "resize"
	InputModePane InputMode = "pane"
	InputModeTab InputMode = "tab"
	InputModeScroll InputMode = "scroll"
	InputModeEnterSearch InputMode = "entersearch"
	InputModeSearch InputMode = "search"
	InputModeRenameTab InputMode = "renametab"
	InputModeSession InputMode = "session"
	InputModeMove InputMode = "move"
	InputModePrompt InputMode = "prompt"
	InputModeTmux InputMode = "tmux"
)

type EventType string

const (
	EventTypeModeUpdate EventType = "ModeUpdate"
	EventTypeTabUpdate EventType = "TabUpdate"
	EventTypePaneUpdate EventType = "PaneUpdate"
	EventTypeKey EventType = "Key"
	EventTypeMouse EventType = "Mouse"
	EventTypeTimer EventType = "Timer"
	EventTypeCopyToClipboard EventType = "CopyToClipboard"
	EventTypeSystemClipboardFailure EventType = "SystemClipboardFailure"
	EventTypeInputReceived EventType = "InputReceived"
	EventTypeVisible EventType = "Visible"
	EventTypeCustomMessage EventType = "CustomMessage"
	EventTypeFileSystemCreate EventType = "FileSystemCreate"
	EventTypeFileSystemRead EventType = "FileSystemRead"
	EventTypeFileSystemUpdate EventType = "FileSystemUpdate"
	EventTypeFileSystemDelete EventType = "FileSystemDelete"
)

