package eiscp

import (
	"log"

	"github.com/colinmurphy1/onkyo-remote/config"
	"github.com/colinmurphy1/onkyo-remote/lib"
)

// Connection struct is stored here for access outside of this package
var Conn *Connection

func init() {
	// Set up custom input names
	noRename := []string{"24", "25", "26", "27", "28", "29", "31", "32", "33", "2B", "2C", "2D", "2E"}
	for _, input := range config.Conf.Inputs {
		// Hex code must be 2 chars in length
		if len(input.Hex) != 2 {
			log.Printf("Invalid hex code length for input \"%s\"\n", input.Name)
		}

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

	// Hide any inputs that are hidden in the yaml configuration file
	EnabledInputs = Inputs
	for _, input := range config.Conf.HiddenInputs {
		// Prevent adding additional inputs that are not in the eiscp spec
		if _, ok := Inputs[input]; !ok {
			log.Println(input, "is not an input code supported by Onkyo. Skipping.")
			continue
		}
		// Remove entry from the EnabledInputs map
		delete(EnabledInputs, input)
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
		//"NRIQSTN", // XML data from receiver (gets model, etc)
	)
}
