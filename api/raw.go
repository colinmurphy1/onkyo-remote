package api

import (
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/help"
	"github.com/gin-gonic/gin"
)

// Set audio source
func SendRaw(c *gin.Context) {
	command := c.Param("command")

	err := eiscp.Conn.SendCmd(command)
	if err != nil {
		help.Response(c, 500, "Error", err)
		return
	}

	help.Response(c, 200, "OK", nil)
}
