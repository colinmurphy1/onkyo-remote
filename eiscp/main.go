package eiscp

import (
	"log"

	"github.com/colinmurphy1/onkyo-remote/config"
)

// Connection struct is stored here for access outside of this package
var Conn *Connection

func init() {
	// Connect to the receiver
	Conn = Onkyo(config.Conf.RECEIVER_IP)

	// Start command watcher goroutine
	go Conn.EiscpWatcher()

	// Initialize OnkyoStatus struct by running QSTN commands on power status,
	// volume level, input, song information, etc...
	// NOTE: It does not matter if you run a QSTN command with the receiver
	// powered on or off, it'll answer with what it has stored.
	startCommands := []string{
		"PWRQSTN", "MVLQSTN", "AMTQSTN", "SLIQSTN", "PRSQSTN", "TUNQSTN",
		"NTIQSTN", "NATQSTN", "NALQSTN", "NTRQSTN",
	}
	for i := 0; i < len(startCommands); i++ {
		err := Conn.SendCmd(startCommands[i])
		if err != nil {
			log.Println("Error running startup command:", err)
		}
	}
}
