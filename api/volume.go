package api

import (
	"strconv"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/help"
	"github.com/gin-gonic/gin"
)

func SetVolume(c *gin.Context) {
	// Convert the volume level from string to int
	volume := c.Param("volume")

	volumeLevel, err := strconv.ParseUint(volume, 10, 32)
	if err != nil {
		// Errors are likely bad data passed, such as a negative volume level.
		help.Response(c, 400, "Bad Request", nil)
		return
	}

	setVolume, _ := eiscp.Conn.SetVolume(uint(volumeLevel))

	res := make(map[string]int)
	res["level"] = int(setVolume)

	help.Response(c, 200, "OK", res)
}

func SetMute(c *gin.Context) {
	status := c.Param("status")

	var set bool

	if status == "on" {
		set = true
	} else if status == "off" {
		set = false
	} else {
		// Invalid option, send HTTP 400
		help.Response(c, 400, "Bad Request", nil)
		return
	}

	muteStatus, err := eiscp.Conn.SetMute(set)

	if err != nil {
		help.Response(c, 500, "Error", err)
		return
	}

	// Make response
	res := make(map[string]bool)
	res["mute"] = muteStatus

	help.Response(c, 200, "OK", res)

}
