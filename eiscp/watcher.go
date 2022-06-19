package eiscp

import (
	"log"
	"strconv"
	"strings"
)

// Watches for responses from the receiver, and configures the status struct accordingly
func (c *Connection) EiscpWatcher() {
	var response, cmd, cmdValue string
	var err error

	for {
		response, err = c.RecvCmd() // Receive command from the receiver
		if err != nil {
			log.Println("RECV ERROR:", err)
			continue // Recover from any errors
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

		// Get input
		case "SLI":
			c.Status.Input.HexCode = cmdValue
			c.Status.Input.Name = Inputs[cmdValue]

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

		// NET service
		case "NMS":
			// NET Source (see data.go for options)
			// This does show other information (such as if you like a song),
			// but this isn't going to be too useful for the controller.
			if c.Status.Input.HexCode == "2B" {
				c.Status.Input.NetSource = NetServices[cmdValue[7:9]]
				continue
			}
			c.Status.Input.NetSource = ""

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
