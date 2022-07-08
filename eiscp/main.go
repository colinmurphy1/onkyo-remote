package eiscp

import (
	"io/ioutil"
	"log"

	"github.com/colinmurphy1/onkyo-remote/config"
	"github.com/colinmurphy1/onkyo-remote/lib"
)

// Connection struct is stored here for access outside of this package
var Conn *Connection

func init() {
	// Disable logging if it is not enabled
	if !config.Conf.Logging.Eiscp {
		log.SetOutput(ioutil.Discard)
	}

	// Set up custom input names
	noRename := []string{"24", "25", "26", "27", "28", "29", "31", "32", "33", "2B", "2C", "2D", "2E"}
	for _, input := range config.Conf.Inputs {
		// Prevent adding additional inputs that are not in the eiscp spec
		if _, ok := Inputs[input.Hex]; !ok {
			log.Println(input.Hex, "is not an input code supported by Onkyo. Skipping.")
			continue
		}

		// Prevent renaming of specific inputs
		if lib.StringInSlice(input.Hex, noRename) {
			log.Println(Inputs[input.Hex], "is not an input that can be renamed. Skipping.")
			continue
		}

		Inputs[input.Hex] = input.Name
	}

	// Connect to the receiver
	Conn = Onkyo(config.Conf.Receiver.Address, config.Conf.Receiver.Port)

	// Set maximum volume in the Onkyo Status struct
	Conn.Status.Volume.Max = config.Conf.MaxVolume

	// Start command watcher goroutine
	go Conn.EiscpWatcher()

	// Initialize OnkyoStatus struct by running QSTN commands.
	// NOTE: It does not matter if you run a QSTN command with the receiver
	// powered on or off, it'll answer with what it has stored.
	Conn.SendMultipleCmds(
		50,
		"PWRQSTN", // Get power status
		"MVLQSTN", // Get volume level
		"AMTQSTN", // Mute status
		"SLIQSTN", // Get Source
		"NJENA",   // Enable album art (jacket)
		"NJALINK", // Send URL instead of raw image data over the wire
	)
}
