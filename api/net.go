package api

import (
	"strings"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

// Send NET playback control commands (play, pause, etc.)
func SetNetPlayback(c *gin.Context) {
	// Valid key options
	keys := []string{
		"PLAY", "STOP", "PAUSE", "TRUP", "TRDN", "REPEAT",
		"RANDOM", "DISPLAY", "RIGHT", "LEFT", "UP",
		"DOWN", "SELECT", "RETURN", "MENU", "TOP", "LIST",
	}

	// Get key parameter and make it uppercase for sending to receiver
	key := c.Param("key")
	key = strings.ToUpper(key)

	// Verify this is a valid key
	if !lib.StringInSlice(key, keys) {
		lib.Response(c, 400, "Invalid NET key", nil)
		return
	}

	// Send command
	err := eiscp.Conn.SendCmd("NTC" + key)
	if err != nil {
		lib.Response(c, 500, "Error", err)
		return
	}

	lib.Response(c, 200, "OK", nil)
}

// Endpoint that updates song title, album, artist, and jacket information
func NetTrackUpdate(c *gin.Context) {
	// Send block of commands
	err := eiscp.Conn.SendMultipleCmds(
		10,
		"NTIQSTN", // Title
		"NALQSTN", // Album
		"NATQSTN", // Artist
		"NJAREQ",  // Album art
	)

	if err != nil {
		lib.Response(c, 500, "Error while sending commands", err)
		return
	}

	lib.Response(c, 200, "OK", nil)
}
