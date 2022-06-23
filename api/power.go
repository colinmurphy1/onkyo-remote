package api

import (
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
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
		lib.Response(c, 400, "Invalid power option", nil)
		return
	}

	err := eiscp.Conn.SetPower(set)

	if err != nil {
		lib.Response(c, 500, "Error", err)
		return
	}

	lib.Response(c, 200, "OK", nil)

}
