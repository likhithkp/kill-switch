package killswitch

import "os"

var Enable bool = false

func init() {
	if !Enable {
		os.Exit(1)
	}
}
