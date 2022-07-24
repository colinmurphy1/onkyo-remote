package api

import (
	"fmt"
	"strings"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

// Set audio source
func SetSource(c *gin.Context) {
	source := c.Param("sourceID")

	// Make source uppercase as the key is uppercase
	source = strings.ToUpper(source)

	// Verify that the source is valid
	if _, ok := eiscp.Conn.Inputs[source]; !ok {
		lib.Response(c, 400, "Invalid input code specified", nil)
		return
	}

	// Send source change command
	if err := eiscp.Conn.SendCmd("SLI" + source); err != nil {
		lib.Response(c, 500, "Error", err)
		return
	}

	lib.Response(c, 200, fmt.Sprintf("Source set to %s", eiscp.Conn.Inputs[source]), nil)
}

// Return source list
func GetSource(c *gin.Context) {
	lib.Response(c, 200, "OK", eiscp.Conn.Inputs)
}
