# go_wakelock
prevent windows from turnning off monitor and sleep

minimum go version **1.18**
# Install
`go get github.com/mhbitarafan/go_wakelock`

# Example
    import wakelock "github.com/mhbitarafan/go_wakelock"
    
    func main() {
    	wakelock.EnableWakeLock() // prevent sleep and turn off monitor
		wakelock.EnableWakeLockSystem() // only prevent sleep
    	wakelock.DisableWakeLock() // release wakelock
    }
    
