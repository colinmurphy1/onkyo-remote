package eiscp

import (
	"io/ioutil"
	"log"

	"github.com/colinmurphy1/onkyo-remote/config"
)

// Connection struct is stored here for access outside of this package
var Conn *Connection

func init() {

	// Disable logging if it is not enabled
	if !config.Conf.EISCP_LOGGING {
		log.SetOutput(ioutil.Discard)
	}

	// Connect to the receiver
	Conn = Onkyo(config.Conf.RECEIVER_IP)

	// Start command watcher goroutine
	go Conn.EiscpWatcher()

	// Initialize OnkyoStatus struct by running QSTN commands.
	// NOTE: It does not matter if you run a QSTN command with the receiver
	// powered on or off, it'll answer with what it has stored.
	startCommands := []string{
		// Get power status
		"PWRQSTN",

		// Get volume level and muting status
		"MVLQSTN",
		"AMTQSTN",

		// Get current source
		"SLIQSTN",

		// Get tuner preset and frequency
		"PRSQSTN",
		"TUNQSTN",

		// Get NET song, album, artist, details
		"NTIQSTN",
		"NATQSTN",
		"NALQSTN",
		"NTRQSTN",

		// Enable album (jacket) art, and have it return a LINK instead of
		// sending a jpeg or bmp image
		"NJAENA",
		"NJALINK",
	}
	for i := 0; i < len(startCommands); i++ {
		err := Conn.SendCmd(startCommands[i])
		if err != nil {
			log.Println("Error running startup command:", err)
		}
	}
}
