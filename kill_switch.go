package killswitch

import "os"

func init() {
	if shouldKill() {
		os.Exit(1)
	}
}

func shouldKill() bool {
	return false
}
