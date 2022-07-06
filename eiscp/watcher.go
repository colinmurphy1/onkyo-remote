package eiscp

import (
	"log"
	"strconv"
	"strings"

	"github.com/colinmurphy1/onkyo-remote/lib"
)

// Watches for responses from the receiver, and configures the status struct accordingly
func (c *Connection) EiscpWatcher() {
	var response, cmd, cmdValue string
	var err error

	for {
		response, err = c.RecvCmd() // Receive command from the receiver
		if err != nil {
			log.Println("RECV ERROR:", err)
			break // Receive errors are generally not recoverable
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
			// Skip sources that do not exist
			if cmdValue == "N/A" {
				continue
			}

			c.Status.Input.HexCode = cmdValue
			c.Status.Input.Name = Inputs[cmdValue]

			// If input is no longer NET (2B), clear some fields
			if c.Status.Input.HexCode != "2B" {
				c.Status.Input.NetSource = ""
				c.Status.SongInfo.Title = ""
				c.Status.SongInfo.Album = ""
				c.Status.SongInfo.Artist = ""
				c.Status.SongInfo.Status = ""
				c.Status.SongInfo.AlbumArt = false
				c.AlbumArt.Data = make([]byte, 0)
				c.AlbumArt.ContentType = ""
			} else {
				// If it is network, send some questions
				c.SendCmd("NMSQSTN")
			}

			// Get information about the source
			c.SendCmd("IFAQSTN")

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
			c.Status.Input.NetSource = NetServices[cmdValue[7:9]]

		// Jacket (Album artwork)
		case "NJA":
			// 2-http://url
			if cmdValue[0:2] == string("2-") {
				// Save album art in memory
				art, ctype, err := lib.GetArt(cmdValue[2:])
				if err != nil {
					log.Println("Error downloading album art:", err)
					continue
				}
				c.AlbumArt.Data = art
				c.AlbumArt.ContentType = ctype    // Content type (eg image/jpeg)
				c.Status.SongInfo.AlbumArt = true // status endpoint reports art is available
			} else if cmdValue == "n-" {
				// Clear out the stored album art
				c.AlbumArt.Data = make([]byte, 0)
				c.AlbumArt.ContentType = ""
				c.Status.SongInfo.AlbumArt = false
			}

		// NET Play status
		case "NST":
			var s string
			playstatus := string(cmdValue[0])
			switch playstatus {
			case "p":
				s = "Paused"
			case "S":
				s = "Stop"
			case "P":
				s = "Play"
			case "F":
				s = "FastForward"
			case "R":
				s = "FastReverse"
			}

			c.Status.SongInfo.Status = s

		// IFA
		case "IFA":
			si := strings.Split(cmdValue, ",")              // Split the output into an array
			c.Status.Input.Info.InputPort = si[0]           // Input port (ANALOG, OPTICAL, COAXIAL?, NETWORK)
			c.Status.Input.Info.InputFormat = si[1]         // Input Format (PCM, DSD?)
			c.Status.Input.Info.SamplingFreq = si[3]        // Unknown
			c.Status.Input.Info.InputSignalChannel = si[4]  // Stereo, Direct, surround sound?, etc.
			c.Status.Input.Info.ListenMode = si[5]          // How many channels
			c.Status.Input.Info.OutputSignalChannel = si[6] // Unknown

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
