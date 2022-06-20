package api

import (
	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/gin-gonic/gin"
)

func GetArt(c *gin.Context) {

	if len(eiscp.Conn.AlbumArt.Data) == 0 {
		c.Data(503, "text/plain", nil)
		return
	}

	// Return the jpeg stored in memory
	c.Data(200, eiscp.Conn.AlbumArt.ContentType, eiscp.Conn.AlbumArt.Data)
}
