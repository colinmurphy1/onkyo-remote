package eiscp

import (
	"log"
	"strconv"
	"strings"

	"github.com/colinmurphy1/onkyo-remote/lib"
)

// Watches for responses from the receiver, and configures the status struct accordingly
func (c *Connection) EiscpWatcher() {
	// Keep track of how many tries to find title, album, and artist information in NET sources.
	var titleRetries, albumRetries, artistRetries int

	// Maximum retries if the title, album, or artist cannot be identified
	var maxRetries int = 5

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
			log.Printf("Power set to %s\n", lib.TOp(pwrStatus, "ON", "OFF"))

		// Get volume level
		case "MVL":
			vol, err := strconv.ParseInt(cmdValue, 16, 64)
			if err != nil {
				log.Println("Could not parse volume level:", err)
				continue // ignore the error and don't continue
			}
			c.Status.Volume.Level = int(vol)
			log.Printf("Volume set to %d\n", vol)

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

			log.Println("Source changed to", c.Status.Input.Name)

			// Reset unneeded status fields on source change
			c.Status.Input.NetSource = ""
			c.Status.SongInfo.Title = ""
			c.Status.SongInfo.Album = ""
			c.Status.SongInfo.Artist = ""
			c.Status.SongInfo.Status = ""
			c.Status.SongInfo.AlbumArt = false
			c.AlbumArt.Data = make([]byte, 0)
			c.AlbumArt.ContentType = ""
			c.Status.Tuner.Frequency = 0
			c.Status.Tuner.Preset = 0

			// On preset type, determine some fields
			switch Conn.Status.Input.HexCode {
			// Network
			case "2B":
				c.SendMultipleCmds(50,
					"NSTQSTN", "NMSQSTN", "NTIQSTN",
					"NATQSTN", "NALQSTN", "NTRQSTN",
				)
			// Tuner
			case "24", "25", "26":
				Conn.SendMultipleCmds(20,
					"PRSQSTN", // Tuner preset
					"TUNQSTN", // Tuner frequency
				)
			}

			// Always get information about the source
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

			// When a track changes, reset the track information
			if ntm[0] == "--:--:--" || ntm[0] == "00:00:01" {
				c.Status.SongInfo.Title = ""
				c.Status.SongInfo.Artist = ""
				c.Status.SongInfo.Album = ""
				titleRetries = 0
				artistRetries = 0
				albumRetries = 0
			}

			// If title is not provided, attempt to get it `maxRetries` times
			if c.Status.SongInfo.Title == "" || c.Status.SongInfo.Title == " " && titleRetries < maxRetries {
				log.Printf("Asking receiver for track title (retry %d)\n", titleRetries)
				c.SendCmd("NTIQSTN")
				titleRetries += 1
			}

			// If album is not provided, attempt to get it `maxRetries` times
			if c.Status.SongInfo.Album == "" || c.Status.SongInfo.Album == " " && albumRetries < maxRetries {
				log.Printf("Asking receiver for album name (retry %d)\n", albumRetries)
				c.SendCmd("NALQSTN")
				albumRetries += 1
			}

			// If artist is not provided, attempt to get it `maxRetries` times
			if c.Status.SongInfo.Artist == "" || c.Status.SongInfo.Artist == " " && artistRetries < maxRetries {
				log.Printf("Asking receiver for artist name (retry %d)\n", artistRetries)
				c.SendCmd("NATQSTN")
				artistRetries += 1
			}

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
			ns := NetServices[cmdValue[7:9]]
			if ns == "" {
				ns = "Unknown"
			}
			c.Status.Input.NetSource = ns
			log.Printf("NETWORK source is %s\n", ns)

		// Jacket (Album artwork)
		case "NJA":
			// 2-http://url
			if cmdValue[0:2] == string("2-") {
				// Download album art off receiver's HTTP server and store in memory
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
