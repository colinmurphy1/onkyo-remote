package api

import (
	"strings"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/help"
	"github.com/gin-gonic/gin"
)

// Set audio source
func SetSource(c *gin.Context) {
	source := c.Param("sourceID")

	// Make source uppercase in the event a lowercase hex value is provided
	source = strings.ToUpper(source)

	// Verify that the source is valid
	if _, ok := eiscp.Inputs[source]; !ok {
		help.Response(c, 400, "Bad Request", "Invalid input, expected HEX code")
		return
	}

	err := eiscp.Conn.SendCmd("SLI" + source)
	if err != nil {
		help.Response(c, 500, "Error", err)
		return
	}

	help.Response(c, 200, "OK", "Source set to "+eiscp.Inputs[source])
}

// Return source list
func GetSource(c *gin.Context) {
	help.Response(c, 200, "OK", eiscp.Inputs)
}
