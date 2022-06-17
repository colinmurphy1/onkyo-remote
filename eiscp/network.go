package eiscp

import (
	"bytes"
	"errors"
	"log"
	"net"
	"strings"
)

// Creates a new Onkyo stereo
func Onkyo(ip string) *Connection {
	device := new(Connection)
	device.ip = ip
	device.iscpVersion = 0x1
	device.iscpDest = 0x31

	// connect
	device.Connect()

	return device
}

// Establish connection to eISCP service. Returns true if connection was successful.
func (c *Connection) Connect() bool {

	log.Printf("Connecting to device at %s", c.ip)

	con, err := net.Dial("tcp", c.ip+":60128")

	// Check for connection errors
	if err != nil {
		log.Fatal("Could not connect to device: ", err)
		c.Disconnect() // Close session
	}

	// Check for proper response
	buffer := make([]byte, 1024)  // Response is stored here
	rlen, err := con.Read(buffer) // Response length
	_ = rlen                      // Don't use the response length

	if err != nil {
		log.Fatal("Could not connect to device: ", err)
		c.Disconnect() // Close session
	}

	// If the response of ISCP is returned, connection is successful
	if string(buffer[:4]) == "ISCP" {
		c.con = con // Store the connection in the struct
		log.Println("Connected!")
		return true
	}
	return false
}

// Close connection to receiver
func (c *Connection) Disconnect() {
	log.Printf("Disconnecting from device at %s", c.ip)
	c.con.Close()
}

// Sends command to receiver. Returns error if unsuccessful
func (c *Connection) SendCmd(command string) error {
	cmd := OnkyoCommand{}
	cmd.Version = c.iscpVersion
	cmd.Destination = c.iscpDest
	cmd.Command = []byte(command)

	// Send command

	log.Println("SEND: ", string(cmd.Command))

	slen, err := c.con.Write(cmd.EiscpCommand())
	_ = slen // We don't care about the response length
	if err != nil {
		c.Disconnect() // Close session
		return err
	}

	return nil
}

// Receives output from receiver
func (c *Connection) RecvCmd() (string, error) {
	// Read repsonse and store it in a 2048 byte buffer
	buffer := make([]byte, 2048)
	_, err := c.con.Read(buffer)

	if err != nil {
		return "", err
	}

	// Verify that the command is valid
	if string(buffer[:4]) != "ISCP" {
		c.Disconnect() // Close session
		return "", errors.New("invalid response from receiver")
	}

	// Split the header and response, giving only the response
	responseSplit := bytes.Split(buffer, []byte{0, 0, 0})[3]

	// Remove the last 3 bytes of the response
	response := string(responseSplit[:len(responseSplit)-3])

	// Responses ending with N/A are invalid
	if strings.HasSuffix(response, "N/A") {
		return response, errors.New("invalid response from receiver")
	}

	log.Println("RECV:", response)

	return response, nil
}
