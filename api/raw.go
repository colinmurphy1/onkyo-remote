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
	if !config.Conf.EnableRaw {
		lib.Response(c, 403, "Endpoint is disabled", nil)
		return
	}

	err := eiscp.Conn.SendCmd(command)
	if err != nil {
		lib.Response(c, 500, "Error", err)
		return
	}

	lib.Response(c, 200, "OK", nil)
}
