package main

import (
	"time"

	wakelock "github.com/mhbitarafan/go_wakelock"
)

func main() {
	println("enabled display and system wakelock")
	wakelock.EnableWakeLock()
	time.Sleep(3 * time.Second)
	println("disabled")
	wakelock.DisableWakeLock()
	time.Sleep(3 * time.Second)
	println("enabled system only wakelock")
	wakelock.EnableWakeLockSystem()
	time.Sleep(3 * time.Second)
}
