package api

import (
	"encoding/base64"

	"github.com/colinmurphy1/onkyo-remote/eiscp"
	"github.com/gin-gonic/gin"
)

// https://png-pixel.com/
var pixel string = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII="

func GetArt(c *gin.Context) {

	// Decode base64 encoded pixel image
	x, _ := base64.StdEncoding.DecodeString(pixel)

	if len(eiscp.Conn.AlbumArt.Data) == 0 {
		c.Data(200, "image/png", x)
		return
	}

	// Return the art stored in memory
	c.Data(200, eiscp.Conn.AlbumArt.ContentType, eiscp.Conn.AlbumArt.Data)
}
