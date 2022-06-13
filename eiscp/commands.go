package eiscp

import (
	"fmt"
	"strings"
)

// Get power status from stereo. Returns bool, true if powered on
func (c *Connection) GetPower() bool {
	res, _ := c.SendCmd("PWRQSTN")
	if res == "!1PWR01" { // 1 is hardcoded; we want the status of the receiver.
		return true
	}
	return false
}

// Turns receiver on or off, returns true if it was successful
func (c *Connection) SetPower(status bool) bool {
	cmd := "PWR00"
	if status {
		cmd = "PWR01"
	}
	// PWR00 = OFF, PWR01 = ON
	_, ok := c.SendCmd(cmd)

	return ok
}

// Sets the volume of the receiver. Returns true if successful
func (c *Connection) SetVolume(vol uint) bool {

	// convert volume to hexadecimal
	volHex := string(fmt.Sprintf("%02x", vol))
	volHex = strings.ToUpper(volHex)

	res, _ := c.SendCmd("MVL" + volHex)

	// Look for proper response
	if res == "!1MVL"+volHex {
		return true
	}
	return false
}

// Get device name
func (c *Connection) GetDeviceName() string {
	panic("Not Implemented")
	//name, _ := c.SendCmd("NDNQSTN")

	//return string(name)
}

// Mute audio. Returns TRUE if successful
func (c *Connection) SetMute(mute bool) bool {
	cmd := "AMT00" // Mute off
	if mute {
		cmd = "AMT01" // Mute on
	}
	_, ok := c.SendCmd(cmd)

	return ok
}

// Returns information about the audio source in a hashmap. If successful, also
// returns true.
func (c *Connection) GetAudioInfo() (map[string]string, bool) {
	res, ok := c.SendCmd("IFAQSTN")

	if !ok {
		return nil, false
	}

	res = res[5:]                 // Remove the !1IFA from the string
	si := strings.Split(res, ",") // Split the output into an array

	sourceInfo := make(map[string]string)
	sourceInfo["InputPort"] = si[0]
	sourceInfo["InputFormat"] = si[1]
	sourceInfo["SamplingFrequency"] = si[3]
	sourceInfo["InputSignalChannel"] = si[4]
	sourceInfo["ListenMode"] = si[5]
	sourceInfo["OutputSignalChannel"] = si[6]

	return sourceInfo, true
}
