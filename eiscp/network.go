package eiscp

import (
	"bytes"
	"errors"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
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

	// Per eISCP spec, allow 50+ ms for a response
	time.Sleep(50 * time.Millisecond)

	return nil
}

// Receives output from receiver
func (c *Connection) RecvCmd() (string, error) {
	// Read repsonse and store it in a 1024 byte buffer
	buffer := make([]byte, 1024)
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

// Watches for responses from the receiver, and configures the status struct accordingly
func (c *Connection) EiscpWatcher() error {
	var response, cmd, cmdValue string
	var err error

	for {
		response, err = c.RecvCmd() // Receive command from the receiver
		if err != nil {
			return err
		}

		cmd = response[2:5]     // iscp command
		cmdValue = response[5:] // iscp command value

		switch cmd {
		// Get power status
		case "PWR":
			pwrStatus := false
			// If the value is 01, the receiver is powered on
			if cmdValue == "01" {
				pwrStatus = true
			}
			c.Status.Power.Status = pwrStatus

		// 12 Volt triggers A, B, and C. Ignored for now.
		// This will be placed in the Power struct, data as 12VA, 12VB, 12VC, etc.
		case "TGA", "TGB", "TGC":
			continue

		// Get volume level
		case "MVL":
			vol, err := strconv.ParseInt(cmdValue, 16, 64)
			if err != nil {
				log.Println("Could not parse volume level:", err)
				continue // ignore the error and don't continue
			}
			c.Status.Volume.Level = int(vol)

		// Get mute status
		case "AMT":
			muteStatus := false

			// If the value is 01, the receiver is muted
			if cmdValue == "01" {
				muteStatus = true
			}

			c.Status.Volume.Mute = muteStatus

		// Get Song Title (NET/USB ONLY)
		case "NTI":
			c.Status.SongInfo.Title = cmdValue

		// Get Artist (NET/USB ONLY)
		case "NAT":
			c.Status.SongInfo.Artist = cmdValue

		// Get Album (NET/USB ONLY)
		case "NAL":
			c.Status.SongInfo.Album = cmdValue

		// Get song time position and length
		case "NTM":
			ntm := strings.Split(cmdValue, "/")
			c.Status.SongInfo.Time.Current = ntm[0]
			c.Status.SongInfo.Time.Length = ntm[1]

		// Track position
		case "NTR":
			ntr := strings.Split(cmdValue, "/")
			c.Status.SongInfo.Track.Current, _ = strconv.Atoi(ntr[0])
			c.Status.SongInfo.Track.Total, _ = strconv.Atoi(ntr[1])

		// Tuner frequency
		case "PRS":
			c.Status.Tuner.Preset, _ = strconv.Atoi(cmdValue)

		// Tuner preset
		case "TUN":
			c.Status.Tuner.Frequency, _ = strconv.ParseFloat(cmdValue, 64)

		// Ignore unknown commands
		default:
			continue
		}
	}
}
