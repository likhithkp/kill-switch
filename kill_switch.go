package killswitch

import "os"

var Enable bool = true

func init() {
	if !Enable {
		os.Exit(1)
	}
}
