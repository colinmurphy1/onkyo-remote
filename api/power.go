package api

import (
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/help"
	"github.com/gin-gonic/gin"
)

func SetPowerStatus(c *gin.Context) {
	status := c.Param("status")

	var set bool

	if status == "on" {
		set = true
	} else if status == "off" {
		set = false
	} else {
		// Invalid option, send HTTP 400
		help.Response(c, 400, "Bad Request", nil)
		return
	}

	err := eiscp.Conn.SetPower(set)

	if err != nil {
		help.Response(c, 500, "Error", err)
		return
	}

	help.Response(c, 200, "OK", nil)

}
