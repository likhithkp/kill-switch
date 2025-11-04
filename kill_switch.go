package killswitch

import "os"

var Enable bool = false

func KillSwitch() {
	if !Enable {
		os.Exit(1)
	}
}
