package api

import (
	"github.com/colinmurphy1/onkyo-remote/help"
	"github.com/gin-gonic/gin"
)

func SetVolume(c *gin.Context) {
	help.Response(c, 501, "Not Implemented", nil)
}

func SetMute(c *gin.Context) {
	help.Response(c, 501, "Not Implemented", nil)
}
