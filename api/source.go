package api

import (
	"strings"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

// Set audio source
func SetSource(c *gin.Context) {
	source := c.Param("sourceID")

	// Make source uppercase in the event a lowercase hex value is provided
	source = strings.ToUpper(source)

	// Verify that the source is valid
	if _, ok := eiscp.Inputs[source]; !ok {
		lib.Response(c, 400, "Invalid input, expected HEX code", nil)
		return
	}

	err := eiscp.Conn.SendCmd("SLI" + source)
	if err != nil {
		lib.Response(c, 500, "Error", err)
		return
	}

	lib.Response(c, 200, "Source set to "+eiscp.Inputs[source], nil)
}

// Return source list
func GetSource(c *gin.Context) {
	lib.Response(c, 200, "OK", eiscp.EnabledInputs)
}
