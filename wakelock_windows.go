package wakelock

import (
	"syscall"
)

const (
	es_continuous       = 0x80000000
	es_system_required  = 0x00000001
	es_display_required = 0x00000002
)

var (
	kernel32                *syscall.DLL
	setThreadExecutionState *syscall.Proc
)

func init() {
	kernel32, _ = syscall.LoadDLL("kernel32.dll")
	defer kernel32.Release()
	setThreadExecutionState = kernel32.MustFindProc("SetThreadExecutionState")
}

func EnableWakeLock() {
	setThreadExecutionState.Call(es_continuous | es_system_required | es_display_required)
}

func EnableWakeLockSystem() {
	setThreadExecutionState.Call(es_continuous | es_system_required)
}

func DisableWakeLock() {
	setThreadExecutionState.Call(es_continuous)
}
