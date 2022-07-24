package api

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/colinmurphy1/onkyo-remote/config"
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

func SetVolume(c *gin.Context) {
	// Convert the volume level from string to int
	volume := c.Param("volume")

	// Convert volume to an unsigned integer
	volumeLevel, err := strconv.ParseUint(volume, 10, 32)
	if err != nil {
		// Errors are likely bad data passed, such as a negative volume level.
		lib.Response(c, 400, "Invalid volume level", nil)
		return
	}

	// Do not allow volume to be set past limit
	if uint(volumeLevel) > config.Conf.MaxVolume {
		lib.Response(c, 400, "Volume exceeds maximum value allowed", nil)
		return
	}

	// convert volume to hexadecimal
	volHex := string(fmt.Sprintf("%02x", volumeLevel))
	volHex = strings.ToUpper(volHex) // must be uppercase per onkyo spec

	// Set volume
	if err := eiscp.Conn.SendCmd("MVL" + volHex); err != nil {
		lib.Response(c, 400, "Could not set volume", nil)
		return
	}

	lib.Response(c, 200, "OK", nil)
}

func SetMute(c *gin.Context) {
	status := c.Param("status")

	// Make status lowercase
	status = strings.ToLower(status)

	var cmd string

	if status == "on" {
		cmd = "AMT01"
	} else if status == "off" {
		cmd = "AMT00"
	} else {
		// Invalid option, send HTTP 400
		lib.Response(c, 400, "Invalid mute option", nil)
		return
	}

	// Send mute command
	if err := eiscp.Conn.SendCmd(cmd); err != nil {
		lib.Response(c, 500, "Could not mute", err)
		return
	}

	lib.Response(c, 200, "OK", nil)
}
