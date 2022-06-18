package main

import (
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
		// STATUS
		routes.GET("/status", api.GetStatus)

		// POWER
		routes.GET("/power/:status", api.SetPowerStatus)

		// VOLUME
		routes.GET("/volume/level/:volume", api.SetVolume)
		routes.GET("/volume/mute/:status", api.SetMute)

		// SOURCE
		routes.GET("/source", api.GetSource)
		routes.GET("/source/:sourceID", api.SetSource)
	}

	// Start http server
	router.Run(":" + config.Conf.HTTP_PORT)
}
