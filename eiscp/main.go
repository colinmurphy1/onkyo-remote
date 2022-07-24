package eiscp

import (
	"github.com/colinmurphy1/onkyo-remote/config"
)

// Connection struct is stored here for access outside of this package
var Conn *Connection

func init() {
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
		100,
		"PWRQSTN", // Get power status
		"MVLQSTN", // Get volume level
		"AMTQSTN", // Mute status
		"NJENA",   // Enable album art (jacket)
		"NJALINK", // Send URL instead of raw image data over the wire
		"NRIQSTN", // XML data from receiver (gets model, etc)
		"SLIQSTN", // Get Source
	)
}
