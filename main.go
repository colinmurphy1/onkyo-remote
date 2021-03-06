package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/colinmurphy1/onkyo-remote/api"
	"github.com/colinmurphy1/onkyo-remote/config"
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/colinmurphy1/onkyo-remote/lib"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var routes *gin.RouterGroup

// Static files
//go:embed static
var s embed.FS

func main() {
	// Disconnect from the receiver when the software terminates
	defer eiscp.Conn.Disconnect()

	// Set up router
	if config.Conf.Logging.HTTP {
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New() // Logs disabled
	}

	// Enable CORS middleware
	router.Use(lib.CORSMiddleware())

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

		// RAW COMMAND (if debug tools are enabled)
		routes.GET("/raw/:command", api.SendRaw)

		// ON-SCREEN DISPLAY
		routes.GET("/osd/:key", api.SetOSD)

		// NET PLAYBACK CONTROLS
		routes.GET("/net/:key", api.SetNetPlayback)

		// NET TRACK UPDATE
		routes.GET("/net/update", api.NetTrackUpdate)

		// TUNER PRESET
		routes.GET("/tuner/preset/:preset", api.SetTunerPreset)

		// ALBUM ART
		routes.GET("/art", api.GetArt)

		// XML DATA (if debug tools are enabled)
		routes.GET("/xml", api.ReceiverXml)
	}

	// Web remote control
	if config.Conf.EnableRemote {
		// Serve static/ as /remote/
		static, _ := fs.Sub(s, "static")
		router.StaticFS("/remote/", http.FS(static))

		// Redirect / to /remote/
		router.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect, "/remote")
		})
	}

	// Start http server
	router.Run(":" + config.Conf.HTTPPort)
}
