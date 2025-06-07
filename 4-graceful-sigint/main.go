//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// stopINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On stopINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If stopINT is called again, just kill the program (last resort)
//

package main

import (
	"os"
	"os/signal"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Create a process
	proc := MockProcess{}

	// Run the process (blocking)
	go proc.Run()

	<-stop
	go proc.Stop()

	signal.Reset(os.Interrupt)

	select {}
}
