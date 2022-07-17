package api

import (
	"github.com/colinmurphy1/onkyo-remote/config"
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

// Send raw ISCP command (for debugging)
func SendRaw(c *gin.Context) {
	command := c.Param("command")

	// check if this endpoint is enabled
	if !config.Conf.EnableDebugTools {
		lib.Response(c, 403, "This endpoint is disabled", nil)
		return
	}

	// Try to send command to receiver
	if err := eiscp.Conn.SendCmd(command); err != nil {
		lib.Response(c, 500, "Error", err)
		return
	}

	lib.Response(c, 200, "OK", nil)
}

// XML data from the receiver
func ReceiverXml(c *gin.Context) {
	// check if this endpoint is enabled
	if !config.Conf.EnableDebugTools {
		c.Data(403, "text/plain", []byte("This endpoint is disabled"))
		return
	}

	// If the data is empty do not return a response
	if len(eiscp.Xml) == 0 {
		c.Data(503, "text/plain", []byte("No XML data to display"))
		return
	}

	// return xml data from receiver
	c.Data(200, "text/xml", []byte(eiscp.Xml))
}
