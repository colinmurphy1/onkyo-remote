package api

import (
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/gin-gonic/gin"
)

func GetArt(c *gin.Context) {

	if len(eiscp.AlbumArt) == 0 {
		c.Data(503, "text/plain", nil)
		return
	}

	// Return the jpeg stored in memory
	c.Data(200, "image/jpeg", eiscp.AlbumArt)
}
