package eiscp

import (
	"bufio"
	"encoding/binary"
	"errors"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/colinmurphy1/onkyo-remote/config"
)

// Creates a new Onkyo stereo
func Onkyo(ip string, port int) *Connection {
	device := new(Connection)
	device.ip = ip
	device.port = port
	device.iscpVersion = 0x1
	device.iscpDest = 0x31

	// connect
	device.Connect()

	return device
}

// Establish connection to eISCP service. Returns true if connection was successful.
func (c *Connection) Connect() bool {
	log.Printf("Connecting to device at %s:%d", c.ip, c.port)

	con, err := net.Dial("tcp", c.ip+":"+strconv.Itoa(c.port))

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
		return false
	}

	c.con = con
	return true
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

	// Print EISCP log if enabled
	if config.Conf.Logging.Eiscp {
		log.Println("SEND:", string(cmd.Command))
	}

	// Send command
	slen, err := c.con.Write(cmd.EiscpCommand())
	_ = slen // We don't care about the response length
	if err != nil {
		c.Disconnect() // Close session
		return err
	}

	return nil
}

// Receives output from receiver

// Credit to reddec on github, used some of his code to figure this out
// https://github.com/reddec/go-eiscp/blob/master/iscp.go
func (c *Connection) RecvCmd() (string, error) {
	reader := bufio.NewReader(c.con)

	// Get the first four bytes header, and verify it says EISCP
	chunk := make([]byte, 4)
	_, err := reader.Read(chunk)
	if err != nil {
		return "", err
	}
	if string(chunk) != "ISCP" {
		return "", errors.New("invalid response from receiver")
	}

	// Get header size
	reader.Read(chunk)
	if binary.BigEndian.Uint32(chunk) != 16 {
		return "", errors.New("invalid header size")
	}

	// Get data size
	reader.Read(chunk)
	dataSize := binary.BigEndian.Uint32(chunk)

	// Skip ISCP version and reserved bytes (iscp version is always 1)
	reserved := make([]byte, 4)
	reader.Read(reserved)

	// Get ISCP response
	iscp := make([]byte, dataSize)
	respLength, _ := reader.Read(iscp)

	// Trim unused bytes
	iscp = iscp[:respLength]

	// If the length of the eISCP response is not the same as the actual response,
	// repeatedly ask for more data until the actual length is equal to the data size
	for len(iscp) != int(dataSize) {
		moreData, err := c.RecvMore()
		if err != nil {
			log.Println("Error receiving extended response from receiver")
			break // give up
		}
		// Append the iscp data to build a full response
		iscp = append(iscp, moreData...)
	}

	// Remove end characters
	response := string(iscp[:len(iscp)-3])

	// Print EISCP log if enabled
	if config.Conf.Logging.Eiscp {
		log.Println("RECV:", response)
	}
	return response, nil
}

// This is essentially the RecvCmd command, without any checks. It is used when
// there is a response that is longer than the
func (c *Connection) RecvMore() ([]byte, error) {
	reader := bufio.NewReader(c.con)

	// Read 2048 bytes
	chunk := make([]byte, 2048)

	respLength, err := reader.Read(chunk)
	if err != nil {
		return nil, errors.New("could not get response from receiver")
	}

	// Trim unused bytes
	response := chunk[:respLength]

	return response, nil
}

// Sends multiple commands in order
// Specify delay in ms as uint, and commands
func (c *Connection) SendMultipleCmds(delay uint, commands ...string) error {
	for _, command := range commands {
		// send command, if there's an error return the error
		err := c.SendCmd(command)
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	return nil
}
