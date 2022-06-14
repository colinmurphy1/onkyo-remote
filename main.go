package main

import (
	_ "net/http"

	"github.com/colinmurphy1/onkyo-remote/api"
	"github.com/colinmurphy1/onkyo-remote/config"
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var routes *gin.RouterGroup

func main() {

	// Disconnect from the receiver when the software terminates
	defer eiscp.Conn.Disconnect()

	// Set up router
	router = gin.Default()

	// Create a router group
	routes = router.Group("/api")
	{
		// POWER
		routes.GET("/power", api.GetPowerStatus)
		routes.GET("/power/set/:status", api.SetPowerStatus)

		// VOLUME
		routes.GET("/volume/set/:volume", api.SetVolume)
	}

	// Start http server
	router.Run(":" + config.Conf.HTTP_PORT)
}
