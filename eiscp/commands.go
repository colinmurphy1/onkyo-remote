package eiscp

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/colinmurphy1/onkyo-remote/config"
)

// Turns receiver on or off
// Returns a boolean with the new power status, and an error if one is present
func (c *Connection) SetPower(status bool) error {
	cmd := "PWR00"
	if status {
		cmd = "PWR01"
	}
	err := c.SendCmd(cmd)

	if err != nil {
		return err
	}

	return nil
}

// Sets volume level
func (c *Connection) SetVolume(vol uint) (int, error) {
	// Ensure volume is less than the limit
	if vol > config.Conf.MaxVolume {
		return 0, errors.New("volume exceeds maximum")
	}

	// convert volume to hexadecimal
	volHex := string(fmt.Sprintf("%02x", vol))
	volHex = strings.ToUpper(volHex) // must be uppercase per onkyo spec

	err := c.SendCmd("MVL" + volHex)

	if err != nil {
		return 0, err
	}

	// Sleep for 200ms to allow command to process
	time.Sleep(200 * time.Millisecond)

	// Return volume level
	return c.Status.Volume.Level, nil
}

// Mute sound, returns new mute status and error if one is present
func (c *Connection) SetMute(mute bool) (bool, error) {
	cmd := "AMT00" // Mute off
	if mute {
		cmd = "AMT01" // Mute on
	}
	err := c.SendCmd(cmd)
	if err != nil {
		return false, err
	}

	// Sleep for 200ms to allow command to process
	time.Sleep(200 * time.Millisecond)

	// Return mute status
	return c.Status.Volume.Mute, nil
}
