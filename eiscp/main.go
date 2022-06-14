package eiscp

import "github.com/colinmurphy1/onkyo-remote/config"

// Connection struct is stored here for access outside of this package
var Conn *Connection

func init() {
	// Connect to the receiver
	Conn = Onkyo(config.Conf.RECEIVER_IP)
}
