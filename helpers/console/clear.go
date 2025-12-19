package console

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

func EnableANSI() {
	if runtime.GOOS != "windows" {
		return
	}

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setConsoleMode := kernel32.NewProc("SetConsoleMode")
	getConsoleMode := kernel32.NewProc("GetConsoleMode")

	stdout := syscall.Handle(os.Stdout.Fd())

	var mode uint32
	getConsoleMode.Call(uintptr(stdout), uintptr(unsafe.Pointer(&mode)))

	const ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004
	setConsoleMode.Call(uintptr(stdout), uintptr(mode|ENABLE_VIRTUAL_TERMINAL_PROCESSING))
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
