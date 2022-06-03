package main

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

// Sets the volume of the receiver. Returns uint with new volume level
func (c *Connection) SetVolume(vol uint) uint {

	// convert volume to hexadecimal
	volHex := string(fmt.Sprintf("%02x", vol))
	volHex = strings.ToUpper(volHex)

	//res, _ := c.SendCmd("MVL" + volHex)

	/*
		if res == "!1MVL"+volHex {

		}
	*/
	return 0
}
