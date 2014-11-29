//
// Keithley S46 RF Switch example
//
//
// The general flow of the code is
//      Open Resource Manager
//      Open VISA Session to an Instrument
//      Write the Identification Query Using viWrite
//      Try to Read a Response With viRead
//      Close the VISA Session
//

package main

import (
	"fmt"

	ke "github.com/jpoirier/visa/keithley"
)

func main() {
	// First we must call viOpenDefaultRM to get the resource manager
	// handle.  We will store this handle in defaultRM.
	rm, status := vi.OpenDefaultRM()
	if status < vi.SUCCESS {
		fmt.Println("Could not open a session to the VISA Resource Manager!")
		return
	}
	defer rm.Close()

	//
	instr, status = ke.OpenGpib(0, 2)
	if status < vi.SUCCESS {
		fmt.Println("An error occurred opening the session to GPIB 0:2")
		return
	}
	defer instr.Close()

	instr.Reset()
	instr.OpenAllChans()
	instr.OpenChan(1)
	instr.CloseChan(1)
	list, _ instr.ClosedChanList()
	fmt.Println("Closed channel list: ", list)
	fmt.Println("Closing Sessions...")
}
