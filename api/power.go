package api

import (
	"strings"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

func SetPowerStatus(c *gin.Context) {
	status := c.Param("status")

	// Make status lowercase
	status = strings.ToLower(status)

	var cmd string

	if status == "on" {
		cmd = "PWR01"
	} else if status == "off" {
		cmd = "PWR00"
	} else {
		lib.Response(c, 400, "Invalid power option", nil)
		return
	}

	// Send command
	if err := eiscp.Conn.SendCmd(cmd); err != nil {
		lib.Response(c, 500, "Error", err)
		return
	}

	lib.Response(c, 200, "OK", nil)

}
