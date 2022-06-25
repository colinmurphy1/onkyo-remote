package api

import (
	"strings"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

// Send OSD (on-screen display) keys
func SetOSD(c *gin.Context) {
	// Valid key options
	keys := []string{
		"MENU", "UP", "DOWN", "LEFT", "RIGHT",
		"ENTER", "EXIT", "AUDIO", "VIDEO", "HOME",
	}

	// Get key parameter and make it uppercase for sending to receiver
	key := c.Param("key")
	key = strings.ToUpper(key)

	// Verify this is a valid key
	if !lib.StringInSlice(key, keys) {
		lib.Response(c, 400, "Invalid OSD key", nil)
		return
	}

	// Send command
	err := eiscp.Conn.SendCmd("OSD" + key)
	if err != nil {
		lib.Response(c, 500, "Error", err)
		return
	}

	lib.Response(c, 200, "OK", nil)
}
