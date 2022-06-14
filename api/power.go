package api

import (
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/help"
	"github.com/gin-gonic/gin"
)

// Get power status of the receiver
func GetPowerStatus(c *gin.Context) {

	res := make(map[string]bool)
	res["status"] = eiscp.Conn.GetPower()

	help.Response(c, 200, "OK", res)
}

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

	// Send power on/off command to receiver. SetPower returns true if it could
	// send the power command to the receiver.
	setPower := eiscp.Conn.SetPower(set)

	if !setPower {
		help.Response(c, 500, "Could not set power status", nil)
	}

	// Generate an api response
	res := make(map[string]bool)
	res["status"] = set
	help.Response(c, 200, "OK", res)
}
