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

	// Set the volume level
	setVolume := eiscp.Conn.SetVolume(uint(volumeLevel))

	if !setVolume {
		// Could not set volume level
		help.Response(c, 500, "Error setting volume level", nil)
		return
	}

	res := make(map[string]int)
	res["level"] = int(volumeLevel)

	help.Response(c, 200, "OK", res)
}
